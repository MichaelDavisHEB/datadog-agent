// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stats.proto

/*
	Package pb is a generated protocol buffer package.

	It is generated from these files:
		stats.proto

	It has these top-level messages:
		StatsPayload
		ClientStatsPayload
		ClientStatsBucket
		ClientGroupedStats
*/
package pb

import (
	fmt "fmt"

	proto "github.com/gogo/protobuf/proto"

	math "math"

	_ "github.com/gogo/protobuf/gogoproto"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// StatsPayload is the payload used to send stats from the agent to the backend.
type StatsPayload struct {
	AgentHostname string               `protobuf:"bytes,1,opt,name=agentHostname,proto3" json:"agentHostname,omitempty"`
	AgentEnv      string               `protobuf:"bytes,2,opt,name=agentEnv,proto3" json:"agentEnv,omitempty"`
	Stats         []ClientStatsPayload `protobuf:"bytes,3,rep,name=stats" json:"stats"`
}

func (m *StatsPayload) Reset()                    { *m = StatsPayload{} }
func (m *StatsPayload) String() string            { return proto.CompactTextString(m) }
func (*StatsPayload) ProtoMessage()               {}
func (*StatsPayload) Descriptor() ([]byte, []int) { return fileDescriptorStats, []int{0} }

func (m *StatsPayload) GetStats() []ClientStatsPayload {
	if m != nil {
		return m.Stats
	}
	return nil
}

// TODO(gbbr): doc
type ClientStatsPayload struct {
	Hostname      string              `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Env           string              `protobuf:"bytes,2,opt,name=env,proto3" json:"env,omitempty"`
	Version       string              `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Stats         []ClientStatsBucket `protobuf:"bytes,4,rep,name=stats" json:"stats"`
	Lang          string              `protobuf:"bytes,5,opt,name=lang,proto3" json:"lang,omitempty"`
	TracerVersion string              `protobuf:"bytes,6,opt,name=tracerVersion,proto3" json:"tracerVersion,omitempty"`
}

func (m *ClientStatsPayload) Reset()                    { *m = ClientStatsPayload{} }
func (m *ClientStatsPayload) String() string            { return proto.CompactTextString(m) }
func (*ClientStatsPayload) ProtoMessage()               {}
func (*ClientStatsPayload) Descriptor() ([]byte, []int) { return fileDescriptorStats, []int{1} }

func (m *ClientStatsPayload) GetStats() []ClientStatsBucket {
	if m != nil {
		return m.Stats
	}
	return nil
}

// TODO(gbbr): doc
type ClientStatsBucket struct {
	Start    uint64               `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	Duration uint64               `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
	Stats    []ClientGroupedStats `protobuf:"bytes,3,rep,name=stats" json:"stats"`
}

func (m *ClientStatsBucket) Reset()                    { *m = ClientStatsBucket{} }
func (m *ClientStatsBucket) String() string            { return proto.CompactTextString(m) }
func (*ClientStatsBucket) ProtoMessage()               {}
func (*ClientStatsBucket) Descriptor() ([]byte, []int) { return fileDescriptorStats, []int{2} }

func (m *ClientStatsBucket) GetStats() []ClientGroupedStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

// TODO(gbbr): doc
type ClientGroupedStats struct {
	Service        string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Resource       string `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	HTTPStatusCode uint32 `protobuf:"varint,4,opt,name=HTTP_status_code,json=HTTPStatusCode,proto3" json:"HTTP_status_code,omitempty"`
	Type           string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	DBType         string `protobuf:"bytes,6,opt,name=DB_type,json=DBType,proto3" json:"DB_type,omitempty"`
	Hits           uint64 `protobuf:"varint,7,opt,name=hits,proto3" json:"hits,omitempty"`
	Errors         uint64 `protobuf:"varint,8,opt,name=errors,proto3" json:"errors,omitempty"`
	Duration       uint64 `protobuf:"varint,9,opt,name=duration,proto3" json:"duration,omitempty"`
	OkSummary      []byte `protobuf:"bytes,10,opt,name=okSummary,proto3" json:"okSummary,omitempty"`
	ErrorSummary   []byte `protobuf:"bytes,11,opt,name=errorSummary,proto3" json:"errorSummary,omitempty"`
	Synthetics     bool   `protobuf:"varint,12,opt,name=synthetics,proto3" json:"synthetics,omitempty"`
	TopLevelHits   uint64 `protobuf:"varint,13,opt,name=topLevelHits,proto3" json:"topLevelHits,omitempty"`
}

func (m *ClientGroupedStats) Reset()                    { *m = ClientGroupedStats{} }
func (m *ClientGroupedStats) String() string            { return proto.CompactTextString(m) }
func (*ClientGroupedStats) ProtoMessage()               {}
func (*ClientGroupedStats) Descriptor() ([]byte, []int) { return fileDescriptorStats, []int{3} }

func init() {
	proto.RegisterType((*StatsPayload)(nil), "pb.StatsPayload")
	proto.RegisterType((*ClientStatsPayload)(nil), "pb.ClientStatsPayload")
	proto.RegisterType((*ClientStatsBucket)(nil), "pb.ClientStatsBucket")
	proto.RegisterType((*ClientGroupedStats)(nil), "pb.ClientGroupedStats")
}
func (m *StatsPayload) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StatsPayload) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.AgentHostname) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintStats(data, i, uint64(len(m.AgentHostname)))
		i += copy(data[i:], m.AgentHostname)
	}
	if len(m.AgentEnv) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintStats(data, i, uint64(len(m.AgentEnv)))
		i += copy(data[i:], m.AgentEnv)
	}
	if len(m.Stats) > 0 {
		for _, msg := range m.Stats {
			data[i] = 0x1a
			i++
			i = encodeVarintStats(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ClientStatsPayload) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ClientStatsPayload) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Hostname) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintStats(data, i, uint64(len(m.Hostname)))
		i += copy(data[i:], m.Hostname)
	}
	if len(m.Env) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintStats(data, i, uint64(len(m.Env)))
		i += copy(data[i:], m.Env)
	}
	if len(m.Version) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintStats(data, i, uint64(len(m.Version)))
		i += copy(data[i:], m.Version)
	}
	if len(m.Stats) > 0 {
		for _, msg := range m.Stats {
			data[i] = 0x22
			i++
			i = encodeVarintStats(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Lang) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintStats(data, i, uint64(len(m.Lang)))
		i += copy(data[i:], m.Lang)
	}
	if len(m.TracerVersion) > 0 {
		data[i] = 0x32
		i++
		i = encodeVarintStats(data, i, uint64(len(m.TracerVersion)))
		i += copy(data[i:], m.TracerVersion)
	}
	return i, nil
}

func (m *ClientStatsBucket) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ClientStatsBucket) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Start != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintStats(data, i, uint64(m.Start))
	}
	if m.Duration != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintStats(data, i, uint64(m.Duration))
	}
	if len(m.Stats) > 0 {
		for _, msg := range m.Stats {
			data[i] = 0x1a
			i++
			i = encodeVarintStats(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ClientGroupedStats) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ClientGroupedStats) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Service) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintStats(data, i, uint64(len(m.Service)))
		i += copy(data[i:], m.Service)
	}
	if len(m.Name) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintStats(data, i, uint64(len(m.Name)))
		i += copy(data[i:], m.Name)
	}
	if len(m.Resource) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintStats(data, i, uint64(len(m.Resource)))
		i += copy(data[i:], m.Resource)
	}
	if m.HTTPStatusCode != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintStats(data, i, uint64(m.HTTPStatusCode))
	}
	if len(m.Type) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintStats(data, i, uint64(len(m.Type)))
		i += copy(data[i:], m.Type)
	}
	if len(m.DBType) > 0 {
		data[i] = 0x32
		i++
		i = encodeVarintStats(data, i, uint64(len(m.DBType)))
		i += copy(data[i:], m.DBType)
	}
	if m.Hits != 0 {
		data[i] = 0x38
		i++
		i = encodeVarintStats(data, i, uint64(m.Hits))
	}
	if m.Errors != 0 {
		data[i] = 0x40
		i++
		i = encodeVarintStats(data, i, uint64(m.Errors))
	}
	if m.Duration != 0 {
		data[i] = 0x48
		i++
		i = encodeVarintStats(data, i, uint64(m.Duration))
	}
	if len(m.OkSummary) > 0 {
		data[i] = 0x52
		i++
		i = encodeVarintStats(data, i, uint64(len(m.OkSummary)))
		i += copy(data[i:], m.OkSummary)
	}
	if len(m.ErrorSummary) > 0 {
		data[i] = 0x5a
		i++
		i = encodeVarintStats(data, i, uint64(len(m.ErrorSummary)))
		i += copy(data[i:], m.ErrorSummary)
	}
	if m.Synthetics {
		data[i] = 0x60
		i++
		if m.Synthetics {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.TopLevelHits != 0 {
		data[i] = 0x68
		i++
		i = encodeVarintStats(data, i, uint64(m.TopLevelHits))
	}
	return i, nil
}

func encodeFixed64Stats(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Stats(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintStats(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *StatsPayload) Size() (n int) {
	var l int
	_ = l
	l = len(m.AgentHostname)
	if l > 0 {
		n += 1 + l + sovStats(uint64(l))
	}
	l = len(m.AgentEnv)
	if l > 0 {
		n += 1 + l + sovStats(uint64(l))
	}
	if len(m.Stats) > 0 {
		for _, e := range m.Stats {
			l = e.Size()
			n += 1 + l + sovStats(uint64(l))
		}
	}
	return n
}

func (m *ClientStatsPayload) Size() (n int) {
	var l int
	_ = l
	l = len(m.Hostname)
	if l > 0 {
		n += 1 + l + sovStats(uint64(l))
	}
	l = len(m.Env)
	if l > 0 {
		n += 1 + l + sovStats(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovStats(uint64(l))
	}
	if len(m.Stats) > 0 {
		for _, e := range m.Stats {
			l = e.Size()
			n += 1 + l + sovStats(uint64(l))
		}
	}
	l = len(m.Lang)
	if l > 0 {
		n += 1 + l + sovStats(uint64(l))
	}
	l = len(m.TracerVersion)
	if l > 0 {
		n += 1 + l + sovStats(uint64(l))
	}
	return n
}

func (m *ClientStatsBucket) Size() (n int) {
	var l int
	_ = l
	if m.Start != 0 {
		n += 1 + sovStats(uint64(m.Start))
	}
	if m.Duration != 0 {
		n += 1 + sovStats(uint64(m.Duration))
	}
	if len(m.Stats) > 0 {
		for _, e := range m.Stats {
			l = e.Size()
			n += 1 + l + sovStats(uint64(l))
		}
	}
	return n
}

func (m *ClientGroupedStats) Size() (n int) {
	var l int
	_ = l
	l = len(m.Service)
	if l > 0 {
		n += 1 + l + sovStats(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovStats(uint64(l))
	}
	l = len(m.Resource)
	if l > 0 {
		n += 1 + l + sovStats(uint64(l))
	}
	if m.HTTPStatusCode != 0 {
		n += 1 + sovStats(uint64(m.HTTPStatusCode))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovStats(uint64(l))
	}
	l = len(m.DBType)
	if l > 0 {
		n += 1 + l + sovStats(uint64(l))
	}
	if m.Hits != 0 {
		n += 1 + sovStats(uint64(m.Hits))
	}
	if m.Errors != 0 {
		n += 1 + sovStats(uint64(m.Errors))
	}
	if m.Duration != 0 {
		n += 1 + sovStats(uint64(m.Duration))
	}
	l = len(m.OkSummary)
	if l > 0 {
		n += 1 + l + sovStats(uint64(l))
	}
	l = len(m.ErrorSummary)
	if l > 0 {
		n += 1 + l + sovStats(uint64(l))
	}
	if m.Synthetics {
		n += 2
	}
	if m.TopLevelHits != 0 {
		n += 1 + sovStats(uint64(m.TopLevelHits))
	}
	return n
}

func sovStats(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStats(x uint64) (n int) {
	return sovStats(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StatsPayload) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatsPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatsPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AgentHostname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AgentHostname = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AgentEnv", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AgentEnv = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stats = append(m.Stats, ClientStatsPayload{})
			if err := m.Stats[len(m.Stats)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStats(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStats
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClientStatsPayload) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClientStatsPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientStatsPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hostname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hostname = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Env", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Env = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stats = append(m.Stats, ClientStatsBucket{})
			if err := m.Stats[len(m.Stats)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lang", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Lang = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TracerVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TracerVersion = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStats(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStats
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClientStatsBucket) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClientStatsBucket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientStatsBucket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			m.Start = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Start |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Duration |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stats = append(m.Stats, ClientGroupedStats{})
			if err := m.Stats[len(m.Stats)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStats(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStats
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClientGroupedStats) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClientGroupedStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientGroupedStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Service", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Service = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resource", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resource = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HTTPStatusCode", wireType)
			}
			m.HTTPStatusCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.HTTPStatusCode |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DBType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DBType = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hits", wireType)
			}
			m.Hits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Hits |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Errors", wireType)
			}
			m.Errors = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Errors |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Duration |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OkSummary", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OkSummary = append(m.OkSummary[:0], data[iNdEx:postIndex]...)
			if m.OkSummary == nil {
				m.OkSummary = []byte{}
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorSummary", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrorSummary = append(m.ErrorSummary[:0], data[iNdEx:postIndex]...)
			if m.ErrorSummary == nil {
				m.ErrorSummary = []byte{}
			}
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Synthetics", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Synthetics = bool(v != 0)
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TopLevelHits", wireType)
			}
			m.TopLevelHits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.TopLevelHits |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStats(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStats
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStats(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStats
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStats
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStats
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthStats
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStats
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipStats(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthStats = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStats   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("stats.proto", fileDescriptorStats) }

var fileDescriptorStats = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xcd, 0x6e, 0xd3, 0x4e,
	0x14, 0xc5, 0x3b, 0x89, 0x9b, 0x8f, 0x9b, 0xe4, 0xaf, 0xfc, 0x47, 0x50, 0x46, 0x15, 0x32, 0x56,
	0xc4, 0xc2, 0x1b, 0x52, 0x51, 0xde, 0xc0, 0x2d, 0x22, 0x0b, 0x16, 0x95, 0x13, 0xb1, 0x8d, 0xfc,
	0x31, 0x38, 0x56, 0x13, 0x8f, 0x35, 0x33, 0x8e, 0xe4, 0x35, 0x3c, 0x00, 0x8f, 0x55, 0x76, 0x3c,
	0x01, 0x42, 0xe1, 0x45, 0xd0, 0x5c, 0xc7, 0x26, 0x6e, 0x25, 0x76, 0x73, 0xce, 0xdc, 0x9c, 0x7b,
	0xe6, 0x17, 0x19, 0x46, 0x4a, 0x07, 0x5a, 0xcd, 0x73, 0x29, 0xb4, 0xa0, 0x9d, 0x3c, 0xbc, 0x7c,
	0x93, 0xa4, 0x7a, 0x53, 0x84, 0xf3, 0x48, 0xec, 0xae, 0x12, 0x91, 0x88, 0x2b, 0xbc, 0x0a, 0x8b,
	0xcf, 0xa8, 0x50, 0xe0, 0xa9, 0xfa, 0xc9, 0xec, 0x2b, 0x81, 0xf1, 0xd2, 0x44, 0xdc, 0x05, 0xe5,
	0x56, 0x04, 0x31, 0x7d, 0x0d, 0x93, 0x20, 0xe1, 0x99, 0x5e, 0x08, 0xa5, 0xb3, 0x60, 0xc7, 0x19,
	0x71, 0x88, 0x3b, 0xf4, 0xdb, 0x26, 0xbd, 0x84, 0x01, 0x1a, 0xef, 0xb3, 0x3d, 0xeb, 0xe0, 0x40,
	0xa3, 0xe9, 0x35, 0x9c, 0x63, 0x29, 0xd6, 0x75, 0xba, 0xee, 0xe8, 0xfa, 0x62, 0x9e, 0x87, 0xf3,
	0x9b, 0x6d, 0xca, 0x33, 0x7d, 0xba, 0xc8, 0xb3, 0x1e, 0x7e, 0xbe, 0x3a, 0xf3, 0xab, 0xd1, 0xd9,
	0x77, 0x02, 0xf4, 0xe9, 0x8c, 0x59, 0xb3, 0x69, 0xf7, 0x68, 0x34, 0x9d, 0x42, 0x97, 0x37, 0xdb,
	0xcd, 0x91, 0x32, 0xe8, 0xef, 0xb9, 0x54, 0xa9, 0xc8, 0x58, 0x17, 0xdd, 0x5a, 0xd2, 0xb7, 0x75,
	0x25, 0x0b, 0x2b, 0x3d, 0x7f, 0x54, 0xc9, 0x2b, 0xa2, 0x7b, 0xae, 0x5b, 0x8d, 0x28, 0x05, 0x6b,
	0x1b, 0x64, 0x09, 0x3b, 0xc7, 0x24, 0x3c, 0x1b, 0x36, 0x5a, 0x06, 0x11, 0x97, 0x9f, 0x8e, 0x6b,
	0x7a, 0x15, 0x9b, 0x96, 0x39, 0x2b, 0xe1, 0xff, 0x27, 0xd9, 0xf4, 0x19, 0x36, 0x90, 0x1a, 0x9f,
	0x61, 0xf9, 0x95, 0x30, 0xef, 0x8b, 0x0b, 0x19, 0x68, 0x93, 0xd5, 0xc1, 0x8b, 0x46, 0xff, 0x03,
	0xe3, 0x07, 0x29, 0x8a, 0x9c, 0xc7, 0x55, 0x7c, 0x0b, 0xe3, 0x97, 0x6e, 0x8d, 0xf1, 0x74, 0xc6,
	0x80, 0x51, 0x5c, 0xee, 0xd3, 0xa8, 0xa6, 0x58, 0x4b, 0xf3, 0x4a, 0x84, 0x5b, 0x51, 0xb4, 0xea,
	0xff, 0x56, 0x72, 0x25, 0x0a, 0x19, 0xf1, 0x23, 0xc7, 0x46, 0x53, 0x17, 0xa6, 0x8b, 0xd5, 0xea,
	0x6e, 0x6d, 0xd6, 0x15, 0x6a, 0x1d, 0x89, 0x98, 0x33, 0xcb, 0x21, 0xee, 0xc4, 0xff, 0xcf, 0xf8,
	0x4b, 0xb4, 0x6f, 0x44, 0x8c, 0xc9, 0xba, 0xcc, 0x79, 0xcd, 0xcf, 0x9c, 0xe9, 0x0b, 0xe8, 0xdf,
	0x7a, 0x6b, 0xb4, 0x2b, 0x72, 0xbd, 0x5b, 0x6f, 0x65, 0x2e, 0x28, 0x58, 0x9b, 0x54, 0x2b, 0xd6,
	0x47, 0x06, 0x78, 0xa6, 0x17, 0xd0, 0xe3, 0x52, 0x0a, 0xa9, 0xd8, 0x00, 0xdd, 0xa3, 0x6a, 0x31,
	0x1b, 0x3e, 0x62, 0xf6, 0x12, 0x86, 0xe2, 0x7e, 0x59, 0xec, 0x76, 0x81, 0x2c, 0x19, 0x38, 0xc4,
	0x1d, 0xfb, 0x7f, 0x0d, 0x3a, 0x83, 0x31, 0x66, 0xd4, 0x03, 0x23, 0x1c, 0x68, 0x79, 0xd4, 0x06,
	0x50, 0x65, 0xa6, 0x37, 0x5c, 0xa7, 0x91, 0x62, 0x63, 0x87, 0xb8, 0x03, 0xff, 0xc4, 0x31, 0x19,
	0x5a, 0xe4, 0x1f, 0xf9, 0x9e, 0x6f, 0x17, 0xa6, 0xf1, 0x04, 0x1b, 0xb4, 0x3c, 0x6f, 0xfa, 0x70,
	0xb0, 0xc9, 0x8f, 0x83, 0x4d, 0x7e, 0x1d, 0x6c, 0xf2, 0xed, 0xb7, 0x7d, 0x16, 0xf6, 0xf0, 0x63,
	0x7b, 0xf7, 0x27, 0x00, 0x00, 0xff, 0xff, 0x36, 0x78, 0x8e, 0x98, 0xae, 0x03, 0x00, 0x00,
}
