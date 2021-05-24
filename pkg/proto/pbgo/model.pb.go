// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datadog/model/v1/model.proto

package pbgo

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EventType int32

const (
	EventType_ADDED    EventType = 0
	EventType_MODIFIED EventType = 1
	EventType_DELETED  EventType = 2
)

var EventType_name = map[int32]string{
	0: "ADDED",
	1: "MODIFIED",
	2: "DELETED",
}

var EventType_value = map[string]int32{
	"ADDED":    0,
	"MODIFIED": 1,
	"DELETED":  2,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bfe05c03006d01ea, []int{0}
}

type TagCardinality int32

const (
	TagCardinality_LOW          TagCardinality = 0
	TagCardinality_ORCHESTRATOR TagCardinality = 1
	TagCardinality_HIGH         TagCardinality = 2
)

var TagCardinality_name = map[int32]string{
	0: "LOW",
	1: "ORCHESTRATOR",
	2: "HIGH",
}

var TagCardinality_value = map[string]int32{
	"LOW":          0,
	"ORCHESTRATOR": 1,
	"HIGH":         2,
}

func (x TagCardinality) String() string {
	return proto.EnumName(TagCardinality_name, int32(x))
}

func (TagCardinality) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bfe05c03006d01ea, []int{1}
}

type HostnameRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HostnameRequest) Reset()         { *m = HostnameRequest{} }
func (m *HostnameRequest) String() string { return proto.CompactTextString(m) }
func (*HostnameRequest) ProtoMessage()    {}
func (*HostnameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe05c03006d01ea, []int{0}
}

func (m *HostnameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostnameRequest.Unmarshal(m, b)
}
func (m *HostnameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostnameRequest.Marshal(b, m, deterministic)
}
func (m *HostnameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostnameRequest.Merge(m, src)
}
func (m *HostnameRequest) XXX_Size() int {
	return xxx_messageInfo_HostnameRequest.Size(m)
}
func (m *HostnameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HostnameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HostnameRequest proto.InternalMessageInfo

// The response message containing the requested hostname
type HostnameReply struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HostnameReply) Reset()         { *m = HostnameReply{} }
func (m *HostnameReply) String() string { return proto.CompactTextString(m) }
func (*HostnameReply) ProtoMessage()    {}
func (*HostnameReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe05c03006d01ea, []int{1}
}

func (m *HostnameReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostnameReply.Unmarshal(m, b)
}
func (m *HostnameReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostnameReply.Marshal(b, m, deterministic)
}
func (m *HostnameReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostnameReply.Merge(m, src)
}
func (m *HostnameReply) XXX_Size() int {
	return xxx_messageInfo_HostnameReply.Size(m)
}
func (m *HostnameReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HostnameReply.DiscardUnknown(m)
}

var xxx_messageInfo_HostnameReply proto.InternalMessageInfo

func (m *HostnameReply) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

type CaptureTriggerRequest struct {
	Duration             string   `protobuf:"bytes,1,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CaptureTriggerRequest) Reset()         { *m = CaptureTriggerRequest{} }
func (m *CaptureTriggerRequest) String() string { return proto.CompactTextString(m) }
func (*CaptureTriggerRequest) ProtoMessage()    {}
func (*CaptureTriggerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe05c03006d01ea, []int{2}
}

func (m *CaptureTriggerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CaptureTriggerRequest.Unmarshal(m, b)
}
func (m *CaptureTriggerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CaptureTriggerRequest.Marshal(b, m, deterministic)
}
func (m *CaptureTriggerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CaptureTriggerRequest.Merge(m, src)
}
func (m *CaptureTriggerRequest) XXX_Size() int {
	return xxx_messageInfo_CaptureTriggerRequest.Size(m)
}
func (m *CaptureTriggerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CaptureTriggerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CaptureTriggerRequest proto.InternalMessageInfo

func (m *CaptureTriggerRequest) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

type CaptureTriggerResponse struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CaptureTriggerResponse) Reset()         { *m = CaptureTriggerResponse{} }
func (m *CaptureTriggerResponse) String() string { return proto.CompactTextString(m) }
func (*CaptureTriggerResponse) ProtoMessage()    {}
func (*CaptureTriggerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe05c03006d01ea, []int{3}
}

func (m *CaptureTriggerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CaptureTriggerResponse.Unmarshal(m, b)
}
func (m *CaptureTriggerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CaptureTriggerResponse.Marshal(b, m, deterministic)
}
func (m *CaptureTriggerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CaptureTriggerResponse.Merge(m, src)
}
func (m *CaptureTriggerResponse) XXX_Size() int {
	return xxx_messageInfo_CaptureTriggerResponse.Size(m)
}
func (m *CaptureTriggerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CaptureTriggerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CaptureTriggerResponse proto.InternalMessageInfo

func (m *CaptureTriggerResponse) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type StreamTagsRequest struct {
	Cardinality          TagCardinality `protobuf:"varint,1,opt,name=cardinality,proto3,enum=datadog.model.v1.TagCardinality" json:"cardinality,omitempty"`
	IncludeFilter        *Filter        `protobuf:"bytes,2,opt,name=includeFilter,proto3" json:"includeFilter,omitempty"`
	ExcludeFilter        *Filter        `protobuf:"bytes,3,opt,name=excludeFilter,proto3" json:"excludeFilter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StreamTagsRequest) Reset()         { *m = StreamTagsRequest{} }
func (m *StreamTagsRequest) String() string { return proto.CompactTextString(m) }
func (*StreamTagsRequest) ProtoMessage()    {}
func (*StreamTagsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe05c03006d01ea, []int{4}
}

func (m *StreamTagsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamTagsRequest.Unmarshal(m, b)
}
func (m *StreamTagsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamTagsRequest.Marshal(b, m, deterministic)
}
func (m *StreamTagsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTagsRequest.Merge(m, src)
}
func (m *StreamTagsRequest) XXX_Size() int {
	return xxx_messageInfo_StreamTagsRequest.Size(m)
}
func (m *StreamTagsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTagsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTagsRequest proto.InternalMessageInfo

func (m *StreamTagsRequest) GetCardinality() TagCardinality {
	if m != nil {
		return m.Cardinality
	}
	return TagCardinality_LOW
}

func (m *StreamTagsRequest) GetIncludeFilter() *Filter {
	if m != nil {
		return m.IncludeFilter
	}
	return nil
}

func (m *StreamTagsRequest) GetExcludeFilter() *Filter {
	if m != nil {
		return m.ExcludeFilter
	}
	return nil
}

type StreamTagsResponse struct {
	Events               []*StreamTagsEvent `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *StreamTagsResponse) Reset()         { *m = StreamTagsResponse{} }
func (m *StreamTagsResponse) String() string { return proto.CompactTextString(m) }
func (*StreamTagsResponse) ProtoMessage()    {}
func (*StreamTagsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe05c03006d01ea, []int{5}
}

func (m *StreamTagsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamTagsResponse.Unmarshal(m, b)
}
func (m *StreamTagsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamTagsResponse.Marshal(b, m, deterministic)
}
func (m *StreamTagsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTagsResponse.Merge(m, src)
}
func (m *StreamTagsResponse) XXX_Size() int {
	return xxx_messageInfo_StreamTagsResponse.Size(m)
}
func (m *StreamTagsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTagsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTagsResponse proto.InternalMessageInfo

func (m *StreamTagsResponse) GetEvents() []*StreamTagsEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

type StreamTagsEvent struct {
	Type                 EventType `protobuf:"varint,1,opt,name=type,proto3,enum=datadog.model.v1.EventType" json:"type,omitempty"`
	Entity               *Entity   `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *StreamTagsEvent) Reset()         { *m = StreamTagsEvent{} }
func (m *StreamTagsEvent) String() string { return proto.CompactTextString(m) }
func (*StreamTagsEvent) ProtoMessage()    {}
func (*StreamTagsEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe05c03006d01ea, []int{6}
}

func (m *StreamTagsEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamTagsEvent.Unmarshal(m, b)
}
func (m *StreamTagsEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamTagsEvent.Marshal(b, m, deterministic)
}
func (m *StreamTagsEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTagsEvent.Merge(m, src)
}
func (m *StreamTagsEvent) XXX_Size() int {
	return xxx_messageInfo_StreamTagsEvent.Size(m)
}
func (m *StreamTagsEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTagsEvent.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTagsEvent proto.InternalMessageInfo

func (m *StreamTagsEvent) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EventType_ADDED
}

func (m *StreamTagsEvent) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

type Filter struct {
	KubeNamespace        string   `protobuf:"bytes,1,opt,name=kubeNamespace,proto3" json:"kubeNamespace,omitempty"`
	Image                string   `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	ContainerName        string   `protobuf:"bytes,3,opt,name=containerName,proto3" json:"containerName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe05c03006d01ea, []int{7}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetKubeNamespace() string {
	if m != nil {
		return m.KubeNamespace
	}
	return ""
}

func (m *Filter) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Filter) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

type Entity struct {
	Id                          *EntityId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Hash                        string    `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	HighCardinalityTags         []string  `protobuf:"bytes,3,rep,name=highCardinalityTags,proto3" json:"highCardinalityTags,omitempty"`
	OrchestratorCardinalityTags []string  `protobuf:"bytes,4,rep,name=orchestratorCardinalityTags,proto3" json:"orchestratorCardinalityTags,omitempty"`
	LowCardinalityTags          []string  `protobuf:"bytes,5,rep,name=lowCardinalityTags,proto3" json:"lowCardinalityTags,omitempty"`
	StandardTags                []string  `protobuf:"bytes,6,rep,name=standardTags,proto3" json:"standardTags,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}  `json:"-"`
	XXX_unrecognized            []byte    `json:"-"`
	XXX_sizecache               int32     `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe05c03006d01ea, []int{8}
}

func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetId() *EntityId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Entity) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Entity) GetHighCardinalityTags() []string {
	if m != nil {
		return m.HighCardinalityTags
	}
	return nil
}

func (m *Entity) GetOrchestratorCardinalityTags() []string {
	if m != nil {
		return m.OrchestratorCardinalityTags
	}
	return nil
}

func (m *Entity) GetLowCardinalityTags() []string {
	if m != nil {
		return m.LowCardinalityTags
	}
	return nil
}

func (m *Entity) GetStandardTags() []string {
	if m != nil {
		return m.StandardTags
	}
	return nil
}

type FetchEntityRequest struct {
	Id                   *EntityId      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Cardinality          TagCardinality `protobuf:"varint,2,opt,name=cardinality,proto3,enum=datadog.model.v1.TagCardinality" json:"cardinality,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FetchEntityRequest) Reset()         { *m = FetchEntityRequest{} }
func (m *FetchEntityRequest) String() string { return proto.CompactTextString(m) }
func (*FetchEntityRequest) ProtoMessage()    {}
func (*FetchEntityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe05c03006d01ea, []int{9}
}

func (m *FetchEntityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchEntityRequest.Unmarshal(m, b)
}
func (m *FetchEntityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchEntityRequest.Marshal(b, m, deterministic)
}
func (m *FetchEntityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchEntityRequest.Merge(m, src)
}
func (m *FetchEntityRequest) XXX_Size() int {
	return xxx_messageInfo_FetchEntityRequest.Size(m)
}
func (m *FetchEntityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchEntityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchEntityRequest proto.InternalMessageInfo

func (m *FetchEntityRequest) GetId() *EntityId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *FetchEntityRequest) GetCardinality() TagCardinality {
	if m != nil {
		return m.Cardinality
	}
	return TagCardinality_LOW
}

type FetchEntityResponse struct {
	Id                   *EntityId      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Cardinality          TagCardinality `protobuf:"varint,2,opt,name=cardinality,proto3,enum=datadog.model.v1.TagCardinality" json:"cardinality,omitempty"`
	Tags                 []string       `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FetchEntityResponse) Reset()         { *m = FetchEntityResponse{} }
func (m *FetchEntityResponse) String() string { return proto.CompactTextString(m) }
func (*FetchEntityResponse) ProtoMessage()    {}
func (*FetchEntityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe05c03006d01ea, []int{10}
}

func (m *FetchEntityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchEntityResponse.Unmarshal(m, b)
}
func (m *FetchEntityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchEntityResponse.Marshal(b, m, deterministic)
}
func (m *FetchEntityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchEntityResponse.Merge(m, src)
}
func (m *FetchEntityResponse) XXX_Size() int {
	return xxx_messageInfo_FetchEntityResponse.Size(m)
}
func (m *FetchEntityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchEntityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchEntityResponse proto.InternalMessageInfo

func (m *FetchEntityResponse) GetId() *EntityId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *FetchEntityResponse) GetCardinality() TagCardinality {
	if m != nil {
		return m.Cardinality
	}
	return TagCardinality_LOW
}

func (m *FetchEntityResponse) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type EntityId struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntityId) Reset()         { *m = EntityId{} }
func (m *EntityId) String() string { return proto.CompactTextString(m) }
func (*EntityId) ProtoMessage()    {}
func (*EntityId) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe05c03006d01ea, []int{11}
}

func (m *EntityId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityId.Unmarshal(m, b)
}
func (m *EntityId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityId.Marshal(b, m, deterministic)
}
func (m *EntityId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityId.Merge(m, src)
}
func (m *EntityId) XXX_Size() int {
	return xxx_messageInfo_EntityId.Size(m)
}
func (m *EntityId) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityId.DiscardUnknown(m)
}

var xxx_messageInfo_EntityId proto.InternalMessageInfo

func (m *EntityId) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *EntityId) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

// UDS Capture
// The message contains the payload and the ancillary info
type UnixDogstatsdMsg struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	PayloadSize          int32    `protobuf:"varint,2,opt,name=payloadSize,proto3" json:"payloadSize,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	AncillarySize        int32    `protobuf:"varint,4,opt,name=ancillarySize,proto3" json:"ancillarySize,omitempty"`
	Ancillary            []byte   `protobuf:"bytes,5,opt,name=ancillary,proto3" json:"ancillary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnixDogstatsdMsg) Reset()         { *m = UnixDogstatsdMsg{} }
func (m *UnixDogstatsdMsg) String() string { return proto.CompactTextString(m) }
func (*UnixDogstatsdMsg) ProtoMessage()    {}
func (*UnixDogstatsdMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe05c03006d01ea, []int{12}
}

func (m *UnixDogstatsdMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnixDogstatsdMsg.Unmarshal(m, b)
}
func (m *UnixDogstatsdMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnixDogstatsdMsg.Marshal(b, m, deterministic)
}
func (m *UnixDogstatsdMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnixDogstatsdMsg.Merge(m, src)
}
func (m *UnixDogstatsdMsg) XXX_Size() int {
	return xxx_messageInfo_UnixDogstatsdMsg.Size(m)
}
func (m *UnixDogstatsdMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_UnixDogstatsdMsg.DiscardUnknown(m)
}

var xxx_messageInfo_UnixDogstatsdMsg proto.InternalMessageInfo

func (m *UnixDogstatsdMsg) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *UnixDogstatsdMsg) GetPayloadSize() int32 {
	if m != nil {
		return m.PayloadSize
	}
	return 0
}

func (m *UnixDogstatsdMsg) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *UnixDogstatsdMsg) GetAncillarySize() int32 {
	if m != nil {
		return m.AncillarySize
	}
	return 0
}

func (m *UnixDogstatsdMsg) GetAncillary() []byte {
	if m != nil {
		return m.Ancillary
	}
	return nil
}

func init() {
	proto.RegisterEnum("datadog.model.v1.EventType", EventType_name, EventType_value)
	proto.RegisterEnum("datadog.model.v1.TagCardinality", TagCardinality_name, TagCardinality_value)
	proto.RegisterType((*HostnameRequest)(nil), "datadog.model.v1.HostnameRequest")
	proto.RegisterType((*HostnameReply)(nil), "datadog.model.v1.HostnameReply")
	proto.RegisterType((*CaptureTriggerRequest)(nil), "datadog.model.v1.CaptureTriggerRequest")
	proto.RegisterType((*CaptureTriggerResponse)(nil), "datadog.model.v1.CaptureTriggerResponse")
	proto.RegisterType((*StreamTagsRequest)(nil), "datadog.model.v1.StreamTagsRequest")
	proto.RegisterType((*StreamTagsResponse)(nil), "datadog.model.v1.StreamTagsResponse")
	proto.RegisterType((*StreamTagsEvent)(nil), "datadog.model.v1.StreamTagsEvent")
	proto.RegisterType((*Filter)(nil), "datadog.model.v1.Filter")
	proto.RegisterType((*Entity)(nil), "datadog.model.v1.Entity")
	proto.RegisterType((*FetchEntityRequest)(nil), "datadog.model.v1.FetchEntityRequest")
	proto.RegisterType((*FetchEntityResponse)(nil), "datadog.model.v1.FetchEntityResponse")
	proto.RegisterType((*EntityId)(nil), "datadog.model.v1.EntityId")
	proto.RegisterType((*UnixDogstatsdMsg)(nil), "datadog.model.v1.UnixDogstatsdMsg")
}

func init() { proto.RegisterFile("datadog/model/v1/model.proto", fileDescriptor_bfe05c03006d01ea) }

var fileDescriptor_bfe05c03006d01ea = []byte{
	// 717 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xdd, 0x4e, 0x13, 0x41,
	0x14, 0x66, 0xfb, 0x47, 0x7b, 0x5a, 0x60, 0x19, 0x94, 0x34, 0xc0, 0x45, 0xdd, 0x78, 0xd1, 0xa0,
	0x69, 0xa1, 0xe8, 0x85, 0x37, 0x46, 0x60, 0x17, 0xdb, 0x04, 0x6c, 0x32, 0xd4, 0x98, 0x78, 0x37,
	0x74, 0xc7, 0xdd, 0x91, 0xfd, 0x73, 0x66, 0x8a, 0xd4, 0x6b, 0x1f, 0xc0, 0x5b, 0xdf, 0xc2, 0xb7,
	0xf1, 0x75, 0xcc, 0x4e, 0xa7, 0xb4, 0x5b, 0x40, 0xa3, 0x17, 0xde, 0x9d, 0xf3, 0x9d, 0xef, 0x3b,
	0x67, 0xf6, 0x9b, 0x33, 0x59, 0xd8, 0x71, 0x89, 0x24, 0x6e, 0xec, 0xb5, 0xc3, 0xd8, 0xa5, 0x41,
	0xfb, 0x6a, 0x7f, 0x12, 0xb4, 0x12, 0x1e, 0xcb, 0x18, 0x99, 0xba, 0xda, 0x9a, 0x80, 0x57, 0xfb,
	0xd6, 0x3a, 0xac, 0x75, 0x63, 0x21, 0x23, 0x12, 0x52, 0x4c, 0x3f, 0x8d, 0xa8, 0x90, 0xd6, 0x13,
	0x58, 0x99, 0x41, 0x49, 0x30, 0x46, 0x5b, 0x50, 0xf6, 0x35, 0x50, 0x37, 0x1a, 0x46, 0xb3, 0x82,
	0x6f, 0x72, 0xeb, 0x00, 0x1e, 0x1e, 0x93, 0x44, 0x8e, 0x38, 0x1d, 0x70, 0xe6, 0x79, 0x94, 0xeb,
	0x2e, 0xa9, 0xc8, 0x1d, 0x71, 0x22, 0x59, 0x1c, 0x4d, 0x45, 0xd3, 0xdc, 0x7a, 0x0a, 0x9b, 0x8b,
	0x22, 0x91, 0xc4, 0x91, 0xa0, 0x08, 0x41, 0x21, 0x21, 0xd2, 0xd7, 0x0a, 0x15, 0x5b, 0x3f, 0x0d,
	0x58, 0x3f, 0x97, 0x9c, 0x92, 0x70, 0x40, 0x3c, 0x31, 0xed, 0x7f, 0x04, 0xd5, 0x21, 0xe1, 0x2e,
	0x8b, 0x48, 0xc0, 0xe4, 0x58, 0x09, 0x56, 0x3b, 0x8d, 0xd6, 0xe2, 0x07, 0xb6, 0x06, 0xc4, 0x3b,
	0x9e, 0xf1, 0xf0, 0xbc, 0x08, 0xbd, 0x84, 0x15, 0x16, 0x0d, 0x83, 0x91, 0x4b, 0x4f, 0x58, 0x20,
	0x29, 0xaf, 0xe7, 0x1a, 0x46, 0xb3, 0xda, 0xa9, 0xdf, 0xee, 0x32, 0xa9, 0xe3, 0x2c, 0x3d, 0xd5,
	0xd3, 0xeb, 0x79, 0x7d, 0xfe, 0x4f, 0xfa, 0x0c, 0xdd, 0xea, 0x03, 0x9a, 0xff, 0x30, 0xed, 0xc1,
	0x0b, 0x28, 0xd1, 0x2b, 0x1a, 0x49, 0x51, 0x37, 0x1a, 0xf9, 0x66, 0xb5, 0xf3, 0xe8, 0x76, 0xbb,
	0x99, 0xca, 0x49, 0x99, 0x58, 0x0b, 0x2c, 0x09, 0x6b, 0x0b, 0x25, 0xd4, 0x86, 0x82, 0x1c, 0x27,
	0x54, 0x1b, 0xb4, 0x7d, 0xbb, 0x97, 0xa2, 0x0d, 0xc6, 0x09, 0xc5, 0x8a, 0x88, 0xf6, 0xa0, 0x44,
	0x23, 0x99, 0x7a, 0x7a, 0xaf, 0x1b, 0x8e, 0xaa, 0x63, 0xcd, 0xb3, 0x3e, 0x42, 0x49, 0x1b, 0xf2,
	0x18, 0x56, 0x2e, 0x47, 0x17, 0xf4, 0x0d, 0x09, 0xa9, 0x48, 0xc8, 0x70, 0xba, 0x2e, 0x59, 0x10,
	0x3d, 0x80, 0x22, 0x0b, 0x89, 0x47, 0xd5, 0x80, 0x0a, 0x9e, 0x24, 0xa9, 0x76, 0x18, 0x47, 0x92,
	0xb0, 0x88, 0xf2, 0x94, 0xab, 0xcc, 0xac, 0xe0, 0x2c, 0x68, 0x7d, 0xcb, 0x41, 0x69, 0x32, 0x1e,
	0xed, 0x42, 0x8e, 0xb9, 0x6a, 0x42, 0xb5, 0xb3, 0x75, 0xdf, 0x21, 0x7b, 0x2e, 0xce, 0x31, 0x37,
	0xdd, 0x2b, 0x9f, 0x08, 0x5f, 0x4f, 0x54, 0x31, 0xda, 0x83, 0x0d, 0x9f, 0x79, 0xfe, 0xdc, 0x76,
	0xa4, 0xae, 0xd5, 0xf3, 0x8d, 0x7c, 0xb3, 0x82, 0xef, 0x2a, 0xa1, 0x57, 0xb0, 0x1d, 0xf3, 0xa1,
	0x4f, 0x85, 0xe4, 0x44, 0xc6, 0x7c, 0x51, 0x59, 0x50, 0xca, 0xdf, 0x51, 0x50, 0x0b, 0x50, 0x10,
	0x7f, 0x5e, 0x14, 0x16, 0x95, 0xf0, 0x8e, 0x0a, 0xb2, 0xa0, 0x26, 0x24, 0x89, 0x5c, 0xc2, 0x5d,
	0xc5, 0x2c, 0x29, 0x66, 0x06, 0xb3, 0xbe, 0x1a, 0x80, 0x4e, 0xa8, 0x1c, 0xfa, 0xfa, 0x5a, 0xf4,
	0x03, 0xf9, 0x1b, 0x7b, 0x16, 0x1e, 0x53, 0xee, 0x1f, 0x1e, 0x93, 0xf5, 0xdd, 0x80, 0x8d, 0xcc,
	0x31, 0xf4, 0x3a, 0xff, 0xe7, 0x73, 0xa4, 0x57, 0x2d, 0x67, 0xf7, 0xa8, 0x62, 0xeb, 0x19, 0x94,
	0xa7, 0x73, 0xd0, 0x26, 0x94, 0x12, 0x4e, 0x3f, 0xb0, 0x6b, 0xbd, 0x9c, 0x3a, 0x43, 0x26, 0xe4,
	0x47, 0xcc, 0xd5, 0x1b, 0x92, 0x86, 0xd6, 0x0f, 0x03, 0xcc, 0xb7, 0x11, 0xbb, 0xb6, 0x63, 0x4f,
	0x48, 0x22, 0x85, 0x7b, 0x26, 0x3c, 0xb4, 0x03, 0x15, 0xc9, 0x42, 0x2a, 0x24, 0x09, 0x13, 0xd5,
	0x21, 0x8f, 0x67, 0x00, 0x6a, 0x40, 0x35, 0x21, 0xe3, 0x20, 0x26, 0xee, 0x39, 0xfb, 0x32, 0x59,
	0xf0, 0x22, 0x9e, 0x87, 0x50, 0x1d, 0x96, 0x75, 0xaa, 0x16, 0xbc, 0x86, 0xa7, 0x69, 0xfa, 0x00,
	0x48, 0x34, 0x64, 0x41, 0x40, 0xf8, 0x58, 0xa9, 0x0b, 0x4a, 0x9d, 0x05, 0xd3, 0xf9, 0x37, 0x40,
	0xbd, 0xa8, 0x3a, 0xcc, 0x80, 0xdd, 0x7d, 0xa8, 0xdc, 0xbc, 0x67, 0x54, 0x81, 0xe2, 0xa1, 0x6d,
	0x3b, 0xb6, 0xb9, 0x84, 0x6a, 0x50, 0x3e, 0xeb, 0xdb, 0xbd, 0x93, 0x9e, 0x63, 0x9b, 0x06, 0xaa,
	0xc2, 0xb2, 0xed, 0x9c, 0x3a, 0x03, 0xc7, 0x36, 0x73, 0xbb, 0xcf, 0x61, 0x35, 0x6b, 0x27, 0x5a,
	0x86, 0xfc, 0x69, 0xff, 0x9d, 0xb9, 0x84, 0x4c, 0xa8, 0xf5, 0xf1, 0x71, 0xd7, 0x39, 0x1f, 0xe0,
	0xc3, 0x41, 0x1f, 0x9b, 0x06, 0x2a, 0x43, 0xa1, 0xdb, 0x7b, 0xdd, 0x35, 0x73, 0x47, 0xe6, 0xfb,
	0xd5, 0xe4, 0xd2, 0x6b, 0xab, 0xff, 0x4a, 0x3b, 0xb9, 0xf0, 0xe2, 0x8b, 0x92, 0x8a, 0x0f, 0x7e,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x62, 0x19, 0x73, 0x83, 0x06, 0x00, 0x00,
}
