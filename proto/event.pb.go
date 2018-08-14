// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package golego_engine_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Event struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Sender               int64             `protobuf:"varint,2,opt,name=sender" json:"sender,omitempty"`
	Meta                 map[string]string `protobuf:"bytes,3,rep,name=meta" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Data                 map[string]string `protobuf:"bytes,4,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Parent               *Event            `protobuf:"bytes,5,opt,name=parent" json:"parent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_2c677269730fc633, []int{0}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetSender() int64 {
	if m != nil {
		return m.Sender
	}
	return 0
}

func (m *Event) GetMeta() map[string]string {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Event) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Event) GetParent() *Event {
	if m != nil {
		return m.Parent
	}
	return nil
}

type EventRequest struct {
	Event                *Event   `protobuf:"bytes,1,opt,name=event" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventRequest) Reset()         { *m = EventRequest{} }
func (m *EventRequest) String() string { return proto.CompactTextString(m) }
func (*EventRequest) ProtoMessage()    {}
func (*EventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_2c677269730fc633, []int{1}
}
func (m *EventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventRequest.Unmarshal(m, b)
}
func (m *EventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventRequest.Marshal(b, m, deterministic)
}
func (dst *EventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRequest.Merge(dst, src)
}
func (m *EventRequest) XXX_Size() int {
	return xxx_messageInfo_EventRequest.Size(m)
}
func (m *EventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventRequest proto.InternalMessageInfo

func (m *EventRequest) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type EventResponse struct {
	Code                 int64             `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message              string            `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Events               []*Event          `protobuf:"bytes,3,rep,name=events" json:"events,omitempty"`
	Results              map[string]string `protobuf:"bytes,4,rep,name=results" json:"results,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *EventResponse) Reset()         { *m = EventResponse{} }
func (m *EventResponse) String() string { return proto.CompactTextString(m) }
func (*EventResponse) ProtoMessage()    {}
func (*EventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_2c677269730fc633, []int{2}
}
func (m *EventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventResponse.Unmarshal(m, b)
}
func (m *EventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventResponse.Marshal(b, m, deterministic)
}
func (dst *EventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventResponse.Merge(dst, src)
}
func (m *EventResponse) XXX_Size() int {
	return xxx_messageInfo_EventResponse.Size(m)
}
func (m *EventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventResponse proto.InternalMessageInfo

func (m *EventResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *EventResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *EventResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *EventResponse) GetResults() map[string]string {
	if m != nil {
		return m.Results
	}
	return nil
}

type RollbackRequest struct {
	Event                *Event            `protobuf:"bytes,1,opt,name=event" json:"event,omitempty"`
	Results              map[string]string `protobuf:"bytes,3,rep,name=results" json:"results,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RollbackRequest) Reset()         { *m = RollbackRequest{} }
func (m *RollbackRequest) String() string { return proto.CompactTextString(m) }
func (*RollbackRequest) ProtoMessage()    {}
func (*RollbackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_2c677269730fc633, []int{3}
}
func (m *RollbackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RollbackRequest.Unmarshal(m, b)
}
func (m *RollbackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RollbackRequest.Marshal(b, m, deterministic)
}
func (dst *RollbackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RollbackRequest.Merge(dst, src)
}
func (m *RollbackRequest) XXX_Size() int {
	return xxx_messageInfo_RollbackRequest.Size(m)
}
func (m *RollbackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RollbackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RollbackRequest proto.InternalMessageInfo

func (m *RollbackRequest) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *RollbackRequest) GetResults() map[string]string {
	if m != nil {
		return m.Results
	}
	return nil
}

type RollbackResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RollbackResponse) Reset()         { *m = RollbackResponse{} }
func (m *RollbackResponse) String() string { return proto.CompactTextString(m) }
func (*RollbackResponse) ProtoMessage()    {}
func (*RollbackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_2c677269730fc633, []int{4}
}
func (m *RollbackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RollbackResponse.Unmarshal(m, b)
}
func (m *RollbackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RollbackResponse.Marshal(b, m, deterministic)
}
func (dst *RollbackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RollbackResponse.Merge(dst, src)
}
func (m *RollbackResponse) XXX_Size() int {
	return xxx_messageInfo_RollbackResponse.Size(m)
}
func (m *RollbackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RollbackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RollbackResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Event)(nil), "golego.engine.proto.Event")
	proto.RegisterMapType((map[string]string)(nil), "golego.engine.proto.Event.DataEntry")
	proto.RegisterMapType((map[string]string)(nil), "golego.engine.proto.Event.MetaEntry")
	proto.RegisterType((*EventRequest)(nil), "golego.engine.proto.EventRequest")
	proto.RegisterType((*EventResponse)(nil), "golego.engine.proto.EventResponse")
	proto.RegisterMapType((map[string]string)(nil), "golego.engine.proto.EventResponse.ResultsEntry")
	proto.RegisterType((*RollbackRequest)(nil), "golego.engine.proto.RollbackRequest")
	proto.RegisterMapType((map[string]string)(nil), "golego.engine.proto.RollbackRequest.ResultsEntry")
	proto.RegisterType((*RollbackResponse)(nil), "golego.engine.proto.RollbackResponse")
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_event_2c677269730fc633) }

var fileDescriptor_event_2c677269730fc633 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xc1, 0x8b, 0xd3, 0x40,
	0x14, 0xc6, 0x9d, 0xa4, 0x69, 0xe9, 0xeb, 0xaa, 0xcb, 0x53, 0x24, 0xe4, 0x14, 0xc3, 0x0a, 0x39,
	0x45, 0x8d, 0x07, 0x97, 0x3d, 0x79, 0x70, 0x0f, 0x22, 0xb2, 0x30, 0x1e, 0x3d, 0x4d, 0x9b, 0x47,
	0x08, 0x4d, 0x27, 0x35, 0x33, 0x2d, 0xf4, 0xaf, 0xf3, 0x2c, 0xf8, 0x07, 0x79, 0x94, 0xcc, 0x4c,
	0x5a, 0x91, 0x9a, 0x5a, 0x61, 0x6f, 0x33, 0x9d, 0xf7, 0xfb, 0xde, 0xf7, 0xbe, 0xd7, 0xc0, 0x8c,
	0xb6, 0x24, 0x75, 0xb6, 0x6e, 0x1b, 0xdd, 0xe0, 0x93, 0xb2, 0xa9, 0xa9, 0x6c, 0x32, 0x92, 0x65,
	0x25, 0xc9, 0xfe, 0x98, 0x7c, 0xf7, 0x20, 0xb8, 0xed, 0x8a, 0xf0, 0x11, 0x78, 0x55, 0x11, 0xb2,
	0x98, 0xa5, 0x53, 0xee, 0x55, 0x05, 0x3e, 0x83, 0xb1, 0x22, 0x59, 0x50, 0x1b, 0x7a, 0x31, 0x4b,
	0x7d, 0xee, 0x6e, 0x78, 0x0d, 0xa3, 0x15, 0x69, 0x11, 0xfa, 0xb1, 0x9f, 0xce, 0xf2, 0xab, 0xec,
	0x88, 0x6a, 0x66, 0x14, 0xb3, 0x4f, 0xa4, 0xc5, 0xad, 0xd4, 0xed, 0x8e, 0x1b, 0xa2, 0x23, 0x0b,
	0xa1, 0x45, 0x38, 0x3a, 0x49, 0xbe, 0x17, 0x7b, 0xb2, 0x23, 0x30, 0x87, 0xf1, 0x5a, 0xb4, 0x24,
	0x75, 0x18, 0xc4, 0x2c, 0x9d, 0xe5, 0xd1, 0xdf, 0x59, 0xee, 0x2a, 0xa3, 0xb7, 0x30, 0xdd, 0x1b,
	0xc0, 0x4b, 0xf0, 0x97, 0xb4, 0x73, 0xd3, 0x75, 0x47, 0x7c, 0x0a, 0xc1, 0x56, 0xd4, 0x1b, 0x32,
	0xd3, 0x4d, 0xb9, 0xbd, 0xdc, 0x78, 0xd7, 0xac, 0x03, 0xf7, 0xfd, 0xcf, 0x01, 0x93, 0x77, 0x70,
	0x61, 0x2d, 0xd0, 0xd7, 0x0d, 0x29, 0x8d, 0xaf, 0x20, 0x30, 0xf9, 0x1b, 0x7a, 0xd8, 0xb4, 0x2d,
	0x4c, 0x7e, 0x32, 0x78, 0xe8, 0x24, 0xd4, 0xba, 0x91, 0x8a, 0x10, 0x61, 0xb4, 0x68, 0x0a, 0x32,
	0x12, 0x3e, 0x37, 0x67, 0x0c, 0x61, 0xb2, 0x22, 0xa5, 0x44, 0xd9, 0x7b, 0xe8, 0xaf, 0x5d, 0x4e,
	0x46, 0x48, 0xb9, 0xed, 0x0c, 0xe6, 0x64, 0x2b, 0xf1, 0x03, 0x4c, 0x5a, 0x52, 0x9b, 0x5a, 0x2b,
	0xb7, 0x98, 0x97, 0x03, 0x90, 0xb3, 0x95, 0x71, 0x4b, 0xd8, 0x1d, 0xf5, 0x7c, 0x74, 0x03, 0x17,
	0xbf, 0x3f, 0x9c, 0x15, 0xde, 0x0f, 0x06, 0x8f, 0x79, 0x53, 0xd7, 0x73, 0xb1, 0x58, 0xfe, 0x77,
	0x80, 0xf8, 0xf1, 0x30, 0x8c, 0x4d, 0xe0, 0xf5, 0x51, 0xe6, 0x8f, 0x46, 0xf7, 0x30, 0x0e, 0xc2,
	0xe5, 0xa1, 0x89, 0x0d, 0x2d, 0xff, 0xc6, 0xdc, 0x1f, 0xe4, 0x33, 0xb5, 0xdb, 0x6a, 0x41, 0xc8,
	0x61, 0x72, 0x27, 0xed, 0xd7, 0xf7, 0x7c, 0x28, 0x74, 0x63, 0x32, 0x4a, 0x4e, 0xef, 0x25, 0x79,
	0x80, 0x5f, 0x00, 0xee, 0x64, 0xdf, 0x1a, 0xaf, 0xfe, 0x65, 0xfc, 0xe8, 0xc5, 0x89, 0xaa, 0x5e,
	0x7c, 0x3e, 0x36, 0x2f, 0x6f, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xb2, 0x67, 0x9a, 0x58,
	0x04, 0x00, 0x00,
}