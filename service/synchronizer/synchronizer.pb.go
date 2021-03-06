// Code generated by protoc-gen-go. DO NOT EDIT.
// source: synchronizer.proto

package gravity_api_synchronizer

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type AssignPipelineRequest struct {
	ClientID             string   `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
	PipelineID           uint64   `protobuf:"varint,2,opt,name=pipelineID,proto3" json:"pipelineID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssignPipelineRequest) Reset()         { *m = AssignPipelineRequest{} }
func (m *AssignPipelineRequest) String() string { return proto.CompactTextString(m) }
func (*AssignPipelineRequest) ProtoMessage()    {}
func (*AssignPipelineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98557a2a36b798c, []int{0}
}

func (m *AssignPipelineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignPipelineRequest.Unmarshal(m, b)
}
func (m *AssignPipelineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignPipelineRequest.Marshal(b, m, deterministic)
}
func (m *AssignPipelineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignPipelineRequest.Merge(m, src)
}
func (m *AssignPipelineRequest) XXX_Size() int {
	return xxx_messageInfo_AssignPipelineRequest.Size(m)
}
func (m *AssignPipelineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignPipelineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AssignPipelineRequest proto.InternalMessageInfo

func (m *AssignPipelineRequest) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *AssignPipelineRequest) GetPipelineID() uint64 {
	if m != nil {
		return m.PipelineID
	}
	return 0
}

type AssignPipelineReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssignPipelineReply) Reset()         { *m = AssignPipelineReply{} }
func (m *AssignPipelineReply) String() string { return proto.CompactTextString(m) }
func (*AssignPipelineReply) ProtoMessage()    {}
func (*AssignPipelineReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98557a2a36b798c, []int{1}
}

func (m *AssignPipelineReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignPipelineReply.Unmarshal(m, b)
}
func (m *AssignPipelineReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignPipelineReply.Marshal(b, m, deterministic)
}
func (m *AssignPipelineReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignPipelineReply.Merge(m, src)
}
func (m *AssignPipelineReply) XXX_Size() int {
	return xxx_messageInfo_AssignPipelineReply.Size(m)
}
func (m *AssignPipelineReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignPipelineReply.DiscardUnknown(m)
}

var xxx_messageInfo_AssignPipelineReply proto.InternalMessageInfo

func (m *AssignPipelineReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *AssignPipelineReply) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type RevokePipelineRequest struct {
	ClientID             string   `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
	PipelineID           uint64   `protobuf:"varint,2,opt,name=pipelineID,proto3" json:"pipelineID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokePipelineRequest) Reset()         { *m = RevokePipelineRequest{} }
func (m *RevokePipelineRequest) String() string { return proto.CompactTextString(m) }
func (*RevokePipelineRequest) ProtoMessage()    {}
func (*RevokePipelineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98557a2a36b798c, []int{2}
}

func (m *RevokePipelineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokePipelineRequest.Unmarshal(m, b)
}
func (m *RevokePipelineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokePipelineRequest.Marshal(b, m, deterministic)
}
func (m *RevokePipelineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokePipelineRequest.Merge(m, src)
}
func (m *RevokePipelineRequest) XXX_Size() int {
	return xxx_messageInfo_RevokePipelineRequest.Size(m)
}
func (m *RevokePipelineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokePipelineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RevokePipelineRequest proto.InternalMessageInfo

func (m *RevokePipelineRequest) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *RevokePipelineRequest) GetPipelineID() uint64 {
	if m != nil {
		return m.PipelineID
	}
	return 0
}

type RevokePipelineReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokePipelineReply) Reset()         { *m = RevokePipelineReply{} }
func (m *RevokePipelineReply) String() string { return proto.CompactTextString(m) }
func (*RevokePipelineReply) ProtoMessage()    {}
func (*RevokePipelineReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f98557a2a36b798c, []int{3}
}

func (m *RevokePipelineReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokePipelineReply.Unmarshal(m, b)
}
func (m *RevokePipelineReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokePipelineReply.Marshal(b, m, deterministic)
}
func (m *RevokePipelineReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokePipelineReply.Merge(m, src)
}
func (m *RevokePipelineReply) XXX_Size() int {
	return xxx_messageInfo_RevokePipelineReply.Size(m)
}
func (m *RevokePipelineReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokePipelineReply.DiscardUnknown(m)
}

var xxx_messageInfo_RevokePipelineReply proto.InternalMessageInfo

func (m *RevokePipelineReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *RevokePipelineReply) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func init() {
	proto.RegisterType((*AssignPipelineRequest)(nil), "gravity.api.synchronizer.AssignPipelineRequest")
	proto.RegisterType((*AssignPipelineReply)(nil), "gravity.api.synchronizer.AssignPipelineReply")
	proto.RegisterType((*RevokePipelineRequest)(nil), "gravity.api.synchronizer.RevokePipelineRequest")
	proto.RegisterType((*RevokePipelineReply)(nil), "gravity.api.synchronizer.RevokePipelineReply")
}

func init() { proto.RegisterFile("synchronizer.proto", fileDescriptor_f98557a2a36b798c) }

var fileDescriptor_f98557a2a36b798c = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0xae, 0xcc, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcb, 0xac, 0x4a, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92,
	0x48, 0x2f, 0x4a, 0x2c, 0xcb, 0x2c, 0xa9, 0xd4, 0x4b, 0x2c, 0xc8, 0xd4, 0x43, 0x96, 0x57, 0x0a,
	0xe6, 0x12, 0x75, 0x2c, 0x2e, 0xce, 0x4c, 0xcf, 0x0b, 0xc8, 0x2c, 0x48, 0xcd, 0xc9, 0xcc, 0x4b,
	0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe2, 0xe2, 0x48, 0xce, 0xc9, 0x4c, 0xcd,
	0x2b, 0xf1, 0x74, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x85, 0xe4, 0xb8, 0xb8,
	0x0a, 0xa0, 0xca, 0x3d, 0x5d, 0x24, 0x98, 0x14, 0x18, 0x35, 0x58, 0x82, 0x90, 0x44, 0x94, 0xdc,
	0xb9, 0x84, 0xd1, 0x0d, 0x2d, 0xc8, 0xa9, 0x14, 0x92, 0xe0, 0x62, 0x2f, 0x2e, 0x4d, 0x4e, 0x4e,
	0x2d, 0x2e, 0x06, 0x9b, 0xc8, 0x11, 0x04, 0xe3, 0x0a, 0x89, 0x71, 0xb1, 0x15, 0xa5, 0x26, 0x16,
	0xe7, 0xe7, 0x81, 0x0d, 0xe3, 0x0c, 0x82, 0xf2, 0x40, 0xae, 0x0b, 0x4a, 0x2d, 0xcb, 0xcf, 0x4e,
	0xa5, 0xb2, 0xeb, 0xd0, 0x0d, 0x25, 0xcb, 0x75, 0x46, 0x7c, 0x5c, 0x3c, 0xc1, 0x48, 0x61, 0x99,
	0xc4, 0x06, 0x0e, 0x6c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x70, 0x43, 0x4c, 0x4e, 0x82,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SynchronizerClient is the client API for Synchronizer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SynchronizerClient interface {
}

type synchronizerClient struct {
	cc *grpc.ClientConn
}

func NewSynchronizerClient(cc *grpc.ClientConn) SynchronizerClient {
	return &synchronizerClient{cc}
}

// SynchronizerServer is the server API for Synchronizer service.
type SynchronizerServer interface {
}

// UnimplementedSynchronizerServer can be embedded to have forward compatible implementations.
type UnimplementedSynchronizerServer struct {
}

func RegisterSynchronizerServer(s *grpc.Server, srv SynchronizerServer) {
	s.RegisterService(&_Synchronizer_serviceDesc, srv)
}

var _Synchronizer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gravity.api.synchronizer.Synchronizer",
	HandlerType: (*SynchronizerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "synchronizer.proto",
}
