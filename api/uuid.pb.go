// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/uuid.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type FetchRequest struct {
	ServiceName          string   `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	ContainerName        string   `protobuf:"bytes,2,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
	NeedCount            int32    `protobuf:"varint,3,opt,name=need_count,json=needCount,proto3" json:"need_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchRequest) Reset()         { *m = FetchRequest{} }
func (m *FetchRequest) String() string { return proto.CompactTextString(m) }
func (*FetchRequest) ProtoMessage()    {}
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61fc83c022ba86aa, []int{0}
}

func (m *FetchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchRequest.Unmarshal(m, b)
}
func (m *FetchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchRequest.Marshal(b, m, deterministic)
}
func (m *FetchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchRequest.Merge(m, src)
}
func (m *FetchRequest) XXX_Size() int {
	return xxx_messageInfo_FetchRequest.Size(m)
}
func (m *FetchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchRequest proto.InternalMessageInfo

func (m *FetchRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *FetchRequest) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

func (m *FetchRequest) GetNeedCount() int32 {
	if m != nil {
		return m.NeedCount
	}
	return 0
}

type UUIDRange struct {
	ServiceId            int32    `protobuf:"varint,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	ContainerId          int32    `protobuf:"varint,2,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	SequenceIdStart      int32    `protobuf:"varint,3,opt,name=sequence_id_start,json=sequenceIdStart,proto3" json:"sequence_id_start,omitempty"`
	SequenceIdEnd        int32    `protobuf:"varint,4,opt,name=sequence_id_end,json=sequenceIdEnd,proto3" json:"sequence_id_end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UUIDRange) Reset()         { *m = UUIDRange{} }
func (m *UUIDRange) String() string { return proto.CompactTextString(m) }
func (*UUIDRange) ProtoMessage()    {}
func (*UUIDRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_61fc83c022ba86aa, []int{1}
}

func (m *UUIDRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UUIDRange.Unmarshal(m, b)
}
func (m *UUIDRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UUIDRange.Marshal(b, m, deterministic)
}
func (m *UUIDRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UUIDRange.Merge(m, src)
}
func (m *UUIDRange) XXX_Size() int {
	return xxx_messageInfo_UUIDRange.Size(m)
}
func (m *UUIDRange) XXX_DiscardUnknown() {
	xxx_messageInfo_UUIDRange.DiscardUnknown(m)
}

var xxx_messageInfo_UUIDRange proto.InternalMessageInfo

func (m *UUIDRange) GetServiceId() int32 {
	if m != nil {
		return m.ServiceId
	}
	return 0
}

func (m *UUIDRange) GetContainerId() int32 {
	if m != nil {
		return m.ContainerId
	}
	return 0
}

func (m *UUIDRange) GetSequenceIdStart() int32 {
	if m != nil {
		return m.SequenceIdStart
	}
	return 0
}

func (m *UUIDRange) GetSequenceIdEnd() int32 {
	if m != nil {
		return m.SequenceIdEnd
	}
	return 0
}

type FetchReply struct {
	Items                []*UUIDRange `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FetchReply) Reset()         { *m = FetchReply{} }
func (m *FetchReply) String() string { return proto.CompactTextString(m) }
func (*FetchReply) ProtoMessage()    {}
func (*FetchReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_61fc83c022ba86aa, []int{2}
}

func (m *FetchReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchReply.Unmarshal(m, b)
}
func (m *FetchReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchReply.Marshal(b, m, deterministic)
}
func (m *FetchReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchReply.Merge(m, src)
}
func (m *FetchReply) XXX_Size() int {
	return xxx_messageInfo_FetchReply.Size(m)
}
func (m *FetchReply) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchReply.DiscardUnknown(m)
}

var xxx_messageInfo_FetchReply proto.InternalMessageInfo

func (m *FetchReply) GetItems() []*UUIDRange {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*FetchRequest)(nil), "api.FetchRequest")
	proto.RegisterType((*UUIDRange)(nil), "api.UUIDRange")
	proto.RegisterType((*FetchReply)(nil), "api.FetchReply")
}

func init() { proto.RegisterFile("api/uuid.proto", fileDescriptor_61fc83c022ba86aa) }

var fileDescriptor_61fc83c022ba86aa = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xdf, 0x4a, 0x84, 0x40,
	0x14, 0xc6, 0xb3, 0x5d, 0x03, 0x8f, 0xbb, 0xca, 0xce, 0x95, 0x04, 0x81, 0x49, 0x85, 0x04, 0x19,
	0x18, 0x3d, 0x41, 0x7f, 0xc0, 0x9b, 0x2e, 0x26, 0xf6, 0x5a, 0x26, 0xe7, 0x50, 0x03, 0xeb, 0x38,
	0xe9, 0x58, 0xec, 0xe3, 0xf4, 0xa6, 0x31, 0xb3, 0xea, 0x7a, 0xfb, 0x3b, 0xbf, 0x99, 0x6f, 0xce,
	0x37, 0x10, 0x30, 0x25, 0xee, 0xfb, 0x5e, 0xf0, 0x4c, 0xb5, 0x8d, 0x6e, 0xc8, 0x82, 0x29, 0x91,
	0xfc, 0xc2, 0xea, 0x15, 0x75, 0xf5, 0x45, 0xf1, 0xbb, 0xc7, 0x4e, 0x93, 0x4b, 0x58, 0x75, 0xd8,
	0xfe, 0x88, 0x0a, 0x4b, 0xc9, 0x6a, 0x8c, 0x9c, 0xd8, 0x49, 0x3d, 0xea, 0x0f, 0xec, 0x8d, 0xd5,
	0x48, 0xae, 0x21, 0xa8, 0x1a, 0xa9, 0x99, 0x90, 0xd8, 0x1e, 0xa4, 0x53, 0x2b, 0xad, 0x27, 0x6a,
	0xb5, 0x0b, 0x00, 0x89, 0xc8, 0xcb, 0xaa, 0xe9, 0xa5, 0x8e, 0x16, 0xb1, 0x93, 0xba, 0xd4, 0x33,
	0xe4, 0xc9, 0x80, 0xe4, 0xcf, 0x01, 0x6f, 0xbb, 0x2d, 0x9e, 0x29, 0x93, 0x9f, 0x56, 0x1e, 0x63,
	0x05, 0xb7, 0xa1, 0x2e, 0xf5, 0x06, 0x52, 0x70, 0xf3, 0xaa, 0x63, 0xa4, 0xe0, 0x36, 0xd0, 0xa5,
	0xfe, 0xc4, 0x0a, 0x4e, 0x6e, 0x61, 0xd3, 0x99, 0x1d, 0xa4, 0xbd, 0xa2, 0xec, 0x34, 0x6b, 0xc7,
	0xd4, 0x70, 0x1c, 0x14, 0xfc, 0xdd, 0x60, 0x72, 0x03, 0xe1, 0xdc, 0x45, 0xc9, 0xa3, 0xa5, 0x35,
	0xd7, 0x47, 0xf3, 0x45, 0xf2, 0x24, 0x07, 0x18, 0xca, 0x51, 0xbb, 0x3d, 0xb9, 0x02, 0x57, 0x68,
	0xac, 0xbb, 0xc8, 0x89, 0x17, 0xa9, 0x9f, 0x07, 0x19, 0x53, 0x22, 0x9b, 0x56, 0xa0, 0x87, 0x61,
	0xfe, 0x08, 0x4b, 0xc3, 0xc8, 0x1d, 0xb8, 0xf6, 0x2c, 0xd9, 0x58, 0x6f, 0x5e, 0xf2, 0x79, 0x38,
	0x47, 0x6a, 0xb7, 0x4f, 0x4e, 0x3e, 0xce, 0xec, 0x9f, 0x3c, 0xfc, 0x07, 0x00, 0x00, 0xff, 0xff,
	0x3f, 0xc9, 0xdb, 0x7b, 0xa5, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UUIDClient is the client API for UUID service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UUIDClient interface {
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchReply, error)
}

type uUIDClient struct {
	cc grpc.ClientConnInterface
}

func NewUUIDClient(cc grpc.ClientConnInterface) UUIDClient {
	return &uUIDClient{cc}
}

func (c *uUIDClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchReply, error) {
	out := new(FetchReply)
	err := c.cc.Invoke(ctx, "/api.UUID/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UUIDServer is the server API for UUID service.
type UUIDServer interface {
	Fetch(context.Context, *FetchRequest) (*FetchReply, error)
}

// UnimplementedUUIDServer can be embedded to have forward compatible implementations.
type UnimplementedUUIDServer struct {
}

func (*UnimplementedUUIDServer) Fetch(ctx context.Context, req *FetchRequest) (*FetchReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}

func RegisterUUIDServer(s *grpc.Server, srv UUIDServer) {
	s.RegisterService(&_UUID_serviceDesc, srv)
}

func _UUID_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UUIDServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UUID/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UUIDServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UUID_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.UUID",
	HandlerType: (*UUIDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _UUID_Fetch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/uuid.proto",
}
