// Code generated by protoc-gen-go. DO NOT EDIT.
// source: random.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Random struct {
	RandomString         string   `protobuf:"bytes,1,opt,name=randomString,proto3" json:"randomString,omitempty"`
	RandomInt            int32    `protobuf:"varint,2,opt,name=randomInt,proto3" json:"randomInt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Random) Reset()         { *m = Random{} }
func (m *Random) String() string { return proto.CompactTextString(m) }
func (*Random) ProtoMessage()    {}
func (*Random) Descriptor() ([]byte, []int) {
	return fileDescriptor_random_2b8f171ebf7b1d7e, []int{0}
}
func (m *Random) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Random.Unmarshal(m, b)
}
func (m *Random) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Random.Marshal(b, m, deterministic)
}
func (dst *Random) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Random.Merge(dst, src)
}
func (m *Random) XXX_Size() int {
	return xxx_messageInfo_Random.Size(m)
}
func (m *Random) XXX_DiscardUnknown() {
	xxx_messageInfo_Random.DiscardUnknown(m)
}

var xxx_messageInfo_Random proto.InternalMessageInfo

func (m *Random) GetRandomString() string {
	if m != nil {
		return m.RandomString
	}
	return ""
}

func (m *Random) GetRandomInt() int32 {
	if m != nil {
		return m.RandomInt
	}
	return 0
}

func init() {
	proto.RegisterType((*Random)(nil), "pb.Random")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RandomServiceClient is the client API for RandomService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RandomServiceClient interface {
	DoSomething(ctx context.Context, in *Random, opts ...grpc.CallOption) (*Random, error)
}

type randomServiceClient struct {
	cc *grpc.ClientConn
}

func NewRandomServiceClient(cc *grpc.ClientConn) RandomServiceClient {
	return &randomServiceClient{cc}
}

func (c *randomServiceClient) DoSomething(ctx context.Context, in *Random, opts ...grpc.CallOption) (*Random, error) {
	out := new(Random)
	err := c.cc.Invoke(ctx, "/pb.RandomService/DoSomething", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RandomServiceServer is the server API for RandomService service.
type RandomServiceServer interface {
	DoSomething(context.Context, *Random) (*Random, error)
}

func RegisterRandomServiceServer(s *grpc.Server, srv RandomServiceServer) {
	s.RegisterService(&_RandomService_serviceDesc, srv)
}

func _RandomService_DoSomething_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Random)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomServiceServer).DoSomething(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RandomService/DoSomething",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomServiceServer).DoSomething(ctx, req.(*Random))
	}
	return interceptor(ctx, in, info, handler)
}

var _RandomService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RandomService",
	HandlerType: (*RandomServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoSomething",
			Handler:    _RandomService_DoSomething_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "random.proto",
}

func init() { proto.RegisterFile("random.proto", fileDescriptor_random_2b8f171ebf7b1d7e) }

var fileDescriptor_random_2b8f171ebf7b1d7e = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4a, 0xcc, 0x4b,
	0xc9, 0xcf, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xf2, 0xe2, 0x62,
	0x0b, 0x02, 0x8b, 0x09, 0x29, 0xc1, 0x64, 0x83, 0x4b, 0x8a, 0x32, 0xf3, 0xd2, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0x50, 0xc4, 0x84, 0x64, 0xb8, 0x38, 0x21, 0x7c, 0xcf, 0xbc, 0x12, 0x09,
	0x26, 0xa0, 0x02, 0xd6, 0x20, 0x84, 0x80, 0x91, 0x05, 0x17, 0x2f, 0xc4, 0xac, 0xe0, 0xd4, 0xa2,
	0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x75, 0x2e, 0x6e, 0x97, 0xfc, 0xe0, 0xfc, 0xdc, 0xd4, 0x92, 0x0c,
	0x90, 0x6e, 0x2e, 0xbd, 0x82, 0x24, 0x3d, 0x88, 0x0a, 0x29, 0x24, 0xb6, 0x12, 0x43, 0x12, 0x1b,
	0xd8, 0x41, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x67, 0xb0, 0xf0, 0x47, 0xa0, 0x00, 0x00,
	0x00,
}
