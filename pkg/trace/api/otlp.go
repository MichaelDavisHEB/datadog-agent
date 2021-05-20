// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package api

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/DataDog/datadog-agent/pkg/trace/config"
	"github.com/DataDog/datadog-agent/pkg/trace/info"
	"github.com/DataDog/datadog-agent/pkg/trace/metrics"
	"github.com/DataDog/datadog-agent/pkg/trace/metrics/timing"
	"github.com/DataDog/datadog-agent/pkg/trace/pb"
	"github.com/DataDog/datadog-agent/pkg/trace/pb/otlppb"
	"github.com/DataDog/datadog-agent/pkg/trace/sampler"
	"github.com/DataDog/datadog-agent/pkg/util/log"

	"github.com/gogo/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	// otlpProtocolHTTP specifies that the incoming connection was made over plain HTTP.
	otlpProtocolHTTP = "http"
	// otlpProtocolGRPC specifies that the incoming connection was made over gRPC.
	otlpProtocolGRPC = "grpc"
)

// OTLPReceiver implements an OpenTelemetry Collector receiver which accepts incoming
// data on two ports for both plain HTTP and gRPC.
type OTLPReceiver struct {
	wg      sync.WaitGroup  // waits for a graceful shutdown
	httpsrv *http.Server    // the running HTTP server on a started receiver, if enabled
	grpcsrv *grpc.Server    // the running GRPC server on a started receiver, if enabled
	out     chan<- *Payload // the outgoing payload channel
	cfg     *config.OTLP    // receiver config
}

// NewOTLPReceiver returns a new OTLPReceiver which sends any incoming traces down the out channel.
func NewOTLPReceiver(out chan<- *Payload, cfg *config.OTLP) *OTLPReceiver {
	return &OTLPReceiver{out: out, cfg: cfg}
}

// Start starts the OTLPReceiver, if any of the servers were configured as active.
func (o *OTLPReceiver) Start() {
	if o.cfg.HTTPPort != 0 {
		o.httpsrv = &http.Server{
			Addr:    fmt.Sprintf("%s:%d", o.cfg.BindHost, o.cfg.HTTPPort),
			Handler: o,
		}
		o.wg.Add(1)
		go func() {
			defer o.wg.Done()
			if err := o.httpsrv.ListenAndServe(); err != nil {
				if err != http.ErrServerClosed {
					log.Criticalf("Error starting OpenTelemetry HTTP server: %v", err)
				}
			}
		}()
		log.Infof("OpenTelemetry HTTP receiver running on http://%s:%d", o.cfg.BindHost, o.cfg.HTTPPort)
	}
	if o.cfg.GRPCPort != 0 {
		ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", o.cfg.BindHost, o.cfg.GRPCPort))
		if err != nil {
			log.Criticalf("Error starting OpenTelemetry gRPC server: %v", err)
		} else {
			o.grpcsrv = grpc.NewServer()
			otlppb.RegisterTraceServiceServer(o.grpcsrv, o)
			o.wg.Add(1)
			go func() {
				defer o.wg.Done()
				if err := o.grpcsrv.Serve(ln); err != nil {
					log.Criticalf("Error starting OpenTelemetry gRPC server: %v", err)
				}
			}()
			log.Infof("OpenTelemetry gRPC receiver running on %s:%d", o.cfg.BindHost, o.cfg.GRPCPort)
		}
	}
}

// Stop stops any running server.
func (o *OTLPReceiver) Stop() {
	if o.httpsrv != nil {
		timeout, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		go func() {
			o.httpsrv.Shutdown(timeout)
			cancel()
		}()
	}
	if o.grpcsrv != nil {
		go o.grpcsrv.Stop()
	}
	o.wg.Wait()
}

// Export implements otlppb.TraceServiceServer
func (o *OTLPReceiver) Export(ctx context.Context, in *otlppb.ExportTraceServiceRequest) (*otlppb.ExportTraceServiceResponse, error) {
	defer timing.Since("datadog.trace_agent.otlp.process_grpc_request_ms", time.Now())
	md, _ := metadata.FromIncomingContext(ctx)
	metrics.Count("datadog.trace_agent.otlp.payload", 1, tagsFromHeaders(http.Header(md), otlpProtocolGRPC), 1)
	o.processRequest(otlpProtocolGRPC, http.Header(md), in)
	return &otlppb.ExportTraceServiceResponse{}, nil
}

// ServeHTTP implements http.Handler
func (o *OTLPReceiver) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer timing.Since("datadog.trace_agent.otlp.process_http_request_ms", time.Now())
	mtags := tagsFromHeaders(req.Header, otlpProtocolHTTP)
	metrics.Count("datadog.trace_agent.otlp.payload", 1, mtags, 1)

	rd := NewLimitedReader(req.Body, o.cfg.MaxRequestBytes)
	slurp, err := ioutil.ReadAll(rd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		metrics.Count("datadog.trace_agent.otlp.error", 1, append(mtags, "reason:read_body"), 1)
		return
	}
	metrics.Count("datadog.trace_agent.otlp.bytes", int64(len(slurp)), mtags, 1)
	var in otlppb.ExportTraceServiceRequest
	switch getMediaType(req) {
	case "application/x-protobuf":
		if err := proto.Unmarshal(slurp, &in); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			metrics.Count("datadog.trace_agent.otlp.error", 1, append(mtags, "reason:decode_proto"), 1)
			return
		}
	case "application/json":
		fallthrough
	default:
		if err := json.Unmarshal(slurp, &in); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			metrics.Count("datadog.trace_agent.otlp.error", 1, append(mtags, "reason:decode_json"), 1)
			return
		}
	}
	o.processRequest(otlpProtocolHTTP, req.Header, &in)
}

func tagsFromHeaders(h http.Header, protocol string) []string {
	return []string{
		"lang:" + h.Get(headerLang),
		"lang_version:" + h.Get(headerLangVersion),
		"interpreter:" + h.Get(headerLangInterpreter),
		"lang_vendor:" + h.Get(headerLangInterpreterVendor),
		"endpoint_version:opentelemetry_" + protocol + "_v1",
	}
}

// processRequest processes the incoming request in. It marks it as received by the given protocol
// using the given headers.
func (o *OTLPReceiver) processRequest(protocol string, header http.Header, in *otlppb.ExportTraceServiceRequest) {
	for _, rspans := range in.ResourceSpans {
		// each rspans is coming from a different resource and should be considered
		// a separate payload; typically there is only one item in this slice
		rattr := make(map[string]string, len(rspans.Resource.Attributes))
		for _, attr := range rspans.Resource.Attributes {
			rattr[attr.Key] = anyValueString(attr.Value)
		}
		lang := rattr["telemetry.sdk.language"]
		if lang == "" {
			lang = header.Get(headerLang)
		}
		tagstats := &info.TagStats{
			Tags: info.Tags{
				Lang:            lang,
				LangVersion:     header.Get(headerLangVersion),
				Interpreter:     header.Get(headerLangInterpreter),
				LangVendor:      header.Get(headerLangInterpreterVendor),
				TracerVersion:   fmt.Sprintf("otlp-%s", rattr["telemetry.sdk.version"]),
				EndpointVersion: fmt.Sprintf("opentelemetry_%s_v1", protocol),
			},
		}
		tracesByID := make(map[uint64]pb.Trace)
		for _, libspans := range rspans.InstrumentationLibrarySpans {
			lib := libspans.InstrumentationLibrary
			for _, span := range libspans.Spans {
				traceID := byteArrayToUint64(span.TraceId)
				if tracesByID[traceID] == nil {
					tracesByID[traceID] = pb.Trace{}
				}
				tracesByID[traceID] = append(tracesByID[traceID], convertSpan(rattr, lib, span))
			}
		}
		metrics.Count("datadog.trace_agent.otlp.spans", int64(len(rspans.InstrumentationLibrarySpans)), tagstats.AsTags(), 1)
		metrics.Count("datadog.trace_agent.otlp.traces", int64(len(tracesByID)), tagstats.AsTags(), 1)
		p := Payload{
			Source:        tagstats,
			ContainerTags: getContainerTags(header.Get(headerContainerID)),
			Traces:        make(pb.Traces, 0, len(tracesByID)),
		}
		for _, trace := range tracesByID {
			p.Traces = append(p.Traces, trace)
		}
		o.out <- &p
	}
}

// convertSpan converts the span in to a Datadog span, and uses the rattr resource tags and the lib instrumentation
// library attributes to further augment it.
func convertSpan(rattr map[string]string, lib *otlppb.InstrumentationLibrary, in *otlppb.Span) *pb.Span {
	name := spanKindName(in.Kind)
	if lib.Name != "" {
		name = lib.Name + "." + name
	}
	span := &pb.Span{
		Name:     name,
		TraceID:  byteArrayToUint64(in.TraceId),
		SpanID:   byteArrayToUint64(in.SpanId),
		ParentID: byteArrayToUint64(in.ParentSpanId),
		Start:    int64(in.StartTimeUnixNano),
		Duration: int64(in.EndTimeUnixNano) - int64(in.StartTimeUnixNano),
		Service:  rattr["service.name"],
		Resource: in.Name,
		Meta:     rattr,
		Metrics: map[string]float64{
			// auto-keep all incoming traces; it was already chosen as a keeper on
			// the client side.
			sampler.KeySamplingPriority: float64(sampler.PriorityAutoKeep),
		},
	}
	// keep original IDs
	span.Meta["otlp_ids.trace"] = hex.EncodeToString(in.TraceId)
	span.Meta["otlp_ids.span"] = hex.EncodeToString(in.SpanId)
	span.Meta["otlp_ids.parent"] = hex.EncodeToString(in.ParentSpanId)
	if _, ok := span.Meta["env"]; !ok {
		if env := span.Meta["deployment.environment"]; env != "" {
			span.Meta["env"] = env
		} else if env := rattr["env"]; env != "" {
			span.Meta["env"] = env
		}
	}
	if _, ok := span.Meta["version"]; !ok {
		if ver := rattr["service.version"]; ver != "" {
			span.Meta["version"] = ver
		}
	}
	if len(in.Events) > 0 {
		out, err := json.Marshal(in.Events)
		if err != nil {
			log.Warnf("Error marshaling span.Event: %v", err)
		} else {
			span.Meta["events"] = string(out)
		}
	}
	for _, kv := range in.Attributes {
		switch v := kv.Value.Value.(type) {
		case *otlppb.AnyValue_DoubleValue:
			span.Metrics[kv.Key] = v.DoubleValue
		case *otlppb.AnyValue_IntValue:
			span.Metrics[kv.Key] = float64(v.IntValue)
		default:
			span.Meta[kv.Key] = anyValueString(kv.Value)
		}
	}
	if in.TraceState != "" {
		span.Meta["trace_state"] = in.TraceState
	}
	if lib.Name != "" {
		span.Meta["instrumentation_library.name"] = lib.Name
	}
	if lib.Version != "" {
		span.Meta["instrumentation_library.version"] = lib.Version
	}
	if span.Service == "" && span.Meta["peer.service"] != "" {
		span.Service = span.Meta["peer.service"]
	}
	if m := span.Meta["http.method"]; m != "" {
		span.Resource = m
		if route := span.Meta["http.route"]; route != "" {
			span.Resource += " " + route
		} else if route := span.Meta["grpc.path"]; route != "" {
			span.Resource += " " + route
		}
	} else if m := span.Meta["messaging.operation"]; m != "" {
		span.Resource = m
		if dest := span.Meta["messaging.destination"]; m != "" {
			span.Resource += " " + dest
		}
	}
	switch in.Kind {
	case otlppb.Span_SPAN_KIND_SERVER:
		span.Type = "web"
	case otlppb.Span_SPAN_KIND_CLIENT:
		span.Type = "http"
		db, ok := span.Meta["db.system"]
		if !ok {
			break
		}
		switch db {
		case "redis", "memcached":
			span.Type = "cache"
		default:
			span.Type = "db"
		}
	default:
		span.Type = "custom"
	}
	if in.Status.Code == otlppb.Status_STATUS_CODE_ERROR {
		span.Error = 1
		for _, e := range in.Events {
			if strings.ToLower(e.Name) != "exception" {
				continue
			}
			for _, attr := range e.Attributes {
				switch attr.Key {
				case "exception.message":
					span.Meta["error.msg"] = anyValueString(attr.Value)
				case "exception.type":
					span.Meta["error.type"] = anyValueString(attr.Value)
				case "exception.stack":
					span.Meta["error.stack"] = anyValueString(attr.Value)
				}
			}
		}
	}
	return span
}

func byteArrayToUint64(b []byte) uint64 {
	if len(b) < 8 {
		return 0
	}
	return binary.LittleEndian.Uint64(b)
}

// anyValueString converts otlppb.AnyValue a to its string representation.
func anyValueString(a *otlppb.AnyValue) string {
	switch v := a.Value.(type) {
	case *otlppb.AnyValue_StringValue:
		return v.StringValue
	case *otlppb.AnyValue_BoolValue:
		if v.BoolValue {
			return "true"
		}
		return "false"
	case *otlppb.AnyValue_IntValue:
		return strconv.FormatInt(v.IntValue, 10)
	case *otlppb.AnyValue_DoubleValue:
		return strconv.FormatFloat(v.DoubleValue, 'f', 2, 64)
	case *otlppb.AnyValue_ArrayValue:
		var str strings.Builder
		for i, val := range v.ArrayValue.Values {
			if i > 0 {
				str.WriteByte(',')
			}
			str.WriteString(anyValueString(val))
		}
		return str.String()
	case *otlppb.AnyValue_KvlistValue:
		var str strings.Builder
		for i, keyval := range v.KvlistValue.Values {
			if i > 0 {
				str.WriteByte(',')
			}
			str.WriteString(keyval.Key)
			str.WriteByte(':')
			str.WriteString(anyValueString(keyval.Value))
		}
		return str.String()
	}
	return a.String() // should never happen
}

var spanKindNames = map[int32]string{
	0: "unspecified",
	1: "internal",
	2: "server",
	3: "client",
	4: "producer",
	5: "consumer",
}

// spanKindName converts the given SpanKind to a valid Datadog span name.
func spanKindName(k otlppb.Span_SpanKind) string {
	name, ok := spanKindNames[int32(k)]
	if !ok {
		return "unknown"
	}
	return name
}