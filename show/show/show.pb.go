// Code generated by protoc-gen-go. DO NOT EDIT.
// source: show.proto

/*
Package lock is a generated protocol buffer package.

It is generated from these files:
	show.proto

It has these top-level messages:
	RequestTime
	ReplyTime
	RequestStr
	ReplyStr
*/
package lock

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

type RequestTime struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *RequestTime) Reset()                    { *m = RequestTime{} }
func (m *RequestTime) String() string            { return proto.CompactTextString(m) }
func (*RequestTime) ProtoMessage()               {}
func (*RequestTime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RequestTime) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReplyTime struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *ReplyTime) Reset()                    { *m = ReplyTime{} }
func (m *ReplyTime) String() string            { return proto.CompactTextString(m) }
func (*ReplyTime) ProtoMessage()               {}
func (*ReplyTime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReplyTime) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type RequestStr struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *RequestStr) Reset()                    { *m = RequestStr{} }
func (m *RequestStr) String() string            { return proto.CompactTextString(m) }
func (*RequestStr) ProtoMessage()               {}
func (*RequestStr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RequestStr) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReplyStr struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *ReplyStr) Reset()                    { *m = ReplyStr{} }
func (m *ReplyStr) String() string            { return proto.CompactTextString(m) }
func (*ReplyStr) ProtoMessage()               {}
func (*ReplyStr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReplyStr) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestTime)(nil), "lock.RequestTime")
	proto.RegisterType((*ReplyTime)(nil), "lock.ReplyTime")
	proto.RegisterType((*RequestStr)(nil), "lock.RequestStr")
	proto.RegisterType((*ReplyStr)(nil), "lock.ReplyStr")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Hello service

type HelloClient interface {
	Lock(ctx context.Context, in *RequestTime, opts ...grpc.CallOption) (Hello_LockClient, error)
	Echo(ctx context.Context, in *RequestStr, opts ...grpc.CallOption) (*ReplyStr, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) Lock(ctx context.Context, in *RequestTime, opts ...grpc.CallOption) (Hello_LockClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Hello_serviceDesc.Streams[0], c.cc, "/lock.Hello/Lock", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloLockClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Hello_LockClient interface {
	Recv() (*ReplyTime, error)
	grpc.ClientStream
}

type helloLockClient struct {
	grpc.ClientStream
}

func (x *helloLockClient) Recv() (*ReplyTime, error) {
	m := new(ReplyTime)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloClient) Echo(ctx context.Context, in *RequestStr, opts ...grpc.CallOption) (*ReplyStr, error) {
	out := new(ReplyStr)
	err := grpc.Invoke(ctx, "/lock.Hello/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hello service

type HelloServer interface {
	Lock(*RequestTime, Hello_LockServer) error
	Echo(context.Context, *RequestStr) (*ReplyStr, error)
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_Lock_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestTime)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServer).Lock(m, &helloLockServer{stream})
}

type Hello_LockServer interface {
	Send(*ReplyTime) error
	grpc.ServerStream
}

type helloLockServer struct {
	grpc.ServerStream
}

func (x *helloLockServer) Send(m *ReplyTime) error {
	return x.ServerStream.SendMsg(m)
}

func _Hello_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestStr)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lock.Hello/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).Echo(ctx, req.(*RequestStr))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lock.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Hello_Echo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Lock",
			Handler:       _Hello_Lock_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "show.proto",
}

func init() { proto.RegisterFile("show.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xce, 0xc8, 0x2f,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0xc9, 0x4f, 0xce, 0x56, 0x52, 0xe7, 0xe2,
	0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x09, 0xc9, 0xcc, 0x4d, 0x15, 0x92, 0xe0, 0x62, 0xcf,
	0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x95,
	0x54, 0xb9, 0x38, 0x83, 0x52, 0x0b, 0x72, 0x2a, 0x09, 0x28, 0x53, 0xe3, 0xe2, 0x82, 0x9a, 0x17,
	0x5c, 0x52, 0x84, 0x47, 0x9d, 0x0a, 0x17, 0x07, 0xd8, 0x38, 0xbc, 0xaa, 0x8c, 0x92, 0xb9, 0x58,
	0x3d, 0x52, 0x73, 0x72, 0xf2, 0x85, 0xf4, 0xb8, 0x58, 0x7c, 0xf2, 0x93, 0xb3, 0x85, 0x04, 0xf5,
	0x40, 0xae, 0xd6, 0x43, 0x72, 0xb2, 0x14, 0x3f, 0x4c, 0x08, 0xea, 0x38, 0x25, 0x06, 0x03, 0x46,
	0x21, 0x2d, 0x2e, 0x16, 0xd7, 0xe4, 0x8c, 0x7c, 0x21, 0x01, 0x14, 0xf5, 0xc1, 0x25, 0x45, 0x52,
	0x7c, 0x48, 0xca, 0x83, 0x4b, 0x8a, 0x94, 0x18, 0x9c, 0x98, 0x02, 0x18, 0x93, 0xd8, 0xc0, 0x61,
	0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x39, 0x63, 0xad, 0xc6, 0x21, 0x01, 0x00, 0x00,
}
