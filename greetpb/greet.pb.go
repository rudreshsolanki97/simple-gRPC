// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greetpb/greet.proto

/*
Package greetpb is a generated protocol buffer package.

It is generated from these files:
	greetpb/greet.proto

It has these top-level messages:
	Greeting
	GreetRequest
	GreetResponse
	GreetManyTimesRequest
	GreetManyTimesResponse
	LongGreetRequest
	LongGreetResponse
	ManyRequest
	ManyResponse
*/
package greetpb

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

type Greeting struct {
	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
}

func (m *Greeting) Reset()                    { *m = Greeting{} }
func (m *Greeting) String() string            { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()               {}
func (*Greeting) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Greeting) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Greeting) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type GreetRequest struct {
	Greeting *Greeting `protobuf:"bytes,1,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *GreetRequest) Reset()                    { *m = GreetRequest{} }
func (m *GreetRequest) String() string            { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()               {}
func (*GreetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GreetRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetResponse struct {
	Result string `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *GreetResponse) Reset()                    { *m = GreetResponse{} }
func (m *GreetResponse) String() string            { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()               {}
func (*GreetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type GreetManyTimesRequest struct {
	Greeting *Greeting `protobuf:"bytes,1,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *GreetManyTimesRequest) Reset()                    { *m = GreetManyTimesRequest{} }
func (m *GreetManyTimesRequest) String() string            { return proto.CompactTextString(m) }
func (*GreetManyTimesRequest) ProtoMessage()               {}
func (*GreetManyTimesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GreetManyTimesRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetManyTimesResponse struct {
	Result string `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *GreetManyTimesResponse) Reset()                    { *m = GreetManyTimesResponse{} }
func (m *GreetManyTimesResponse) String() string            { return proto.CompactTextString(m) }
func (*GreetManyTimesResponse) ProtoMessage()               {}
func (*GreetManyTimesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GreetManyTimesResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type LongGreetRequest struct {
	Greeting *Greeting `protobuf:"bytes,1,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *LongGreetRequest) Reset()                    { *m = LongGreetRequest{} }
func (m *LongGreetRequest) String() string            { return proto.CompactTextString(m) }
func (*LongGreetRequest) ProtoMessage()               {}
func (*LongGreetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LongGreetRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type LongGreetResponse struct {
	Result string `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *LongGreetResponse) Reset()                    { *m = LongGreetResponse{} }
func (m *LongGreetResponse) String() string            { return proto.CompactTextString(m) }
func (*LongGreetResponse) ProtoMessage()               {}
func (*LongGreetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *LongGreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type ManyRequest struct {
	Greeting *Greeting `protobuf:"bytes,1,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *ManyRequest) Reset()                    { *m = ManyRequest{} }
func (m *ManyRequest) String() string            { return proto.CompactTextString(m) }
func (*ManyRequest) ProtoMessage()               {}
func (*ManyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ManyRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type ManyResponse struct {
	Result string `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *ManyResponse) Reset()                    { *m = ManyResponse{} }
func (m *ManyResponse) String() string            { return proto.CompactTextString(m) }
func (*ManyResponse) ProtoMessage()               {}
func (*ManyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ManyResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeting)(nil), "greet.Greeting")
	proto.RegisterType((*GreetRequest)(nil), "greet.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "greet.GreetResponse")
	proto.RegisterType((*GreetManyTimesRequest)(nil), "greet.GreetManyTimesRequest")
	proto.RegisterType((*GreetManyTimesResponse)(nil), "greet.GreetManyTimesResponse")
	proto.RegisterType((*LongGreetRequest)(nil), "greet.LongGreetRequest")
	proto.RegisterType((*LongGreetResponse)(nil), "greet.LongGreetResponse")
	proto.RegisterType((*ManyRequest)(nil), "greet.ManyRequest")
	proto.RegisterType((*ManyResponse)(nil), "greet.ManyResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GreetService service

type GreetServiceClient interface {
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	GreetManyTimes(ctx context.Context, in *GreetManyTimesRequest, opts ...grpc.CallOption) (GreetService_GreetManyTimesClient, error)
	LongGreet(ctx context.Context, opts ...grpc.CallOption) (GreetService_LongGreetClient, error)
	BiDiGreet(ctx context.Context, opts ...grpc.CallOption) (GreetService_BiDiGreetClient, error)
}

type greetServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreetServiceClient(cc *grpc.ClientConn) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := grpc.Invoke(ctx, "/greet.GreetService/Greet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) GreetManyTimes(ctx context.Context, in *GreetManyTimesRequest, opts ...grpc.CallOption) (GreetService_GreetManyTimesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GreetService_serviceDesc.Streams[0], c.cc, "/greet.GreetService/GreetManyTimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetManyTimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_GreetManyTimesClient interface {
	Recv() (*GreetManyTimesResponse, error)
	grpc.ClientStream
}

type greetServiceGreetManyTimesClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetManyTimesClient) Recv() (*GreetManyTimesResponse, error) {
	m := new(GreetManyTimesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) LongGreet(ctx context.Context, opts ...grpc.CallOption) (GreetService_LongGreetClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GreetService_serviceDesc.Streams[1], c.cc, "/greet.GreetService/LongGreet", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceLongGreetClient{stream}
	return x, nil
}

type GreetService_LongGreetClient interface {
	Send(*LongGreetRequest) error
	CloseAndRecv() (*LongGreetResponse, error)
	grpc.ClientStream
}

type greetServiceLongGreetClient struct {
	grpc.ClientStream
}

func (x *greetServiceLongGreetClient) Send(m *LongGreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceLongGreetClient) CloseAndRecv() (*LongGreetResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(LongGreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) BiDiGreet(ctx context.Context, opts ...grpc.CallOption) (GreetService_BiDiGreetClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GreetService_serviceDesc.Streams[2], c.cc, "/greet.GreetService/BiDiGreet", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceBiDiGreetClient{stream}
	return x, nil
}

type GreetService_BiDiGreetClient interface {
	Send(*ManyRequest) error
	Recv() (*ManyResponse, error)
	grpc.ClientStream
}

type greetServiceBiDiGreetClient struct {
	grpc.ClientStream
}

func (x *greetServiceBiDiGreetClient) Send(m *ManyRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceBiDiGreetClient) Recv() (*ManyResponse, error) {
	m := new(ManyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GreetService service

type GreetServiceServer interface {
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
	GreetManyTimes(*GreetManyTimesRequest, GreetService_GreetManyTimesServer) error
	LongGreet(GreetService_LongGreetServer) error
	BiDiGreet(GreetService_BiDiGreetServer) error
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_GreetManyTimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetManyTimesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).GreetManyTimes(m, &greetServiceGreetManyTimesServer{stream})
}

type GreetService_GreetManyTimesServer interface {
	Send(*GreetManyTimesResponse) error
	grpc.ServerStream
}

type greetServiceGreetManyTimesServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetManyTimesServer) Send(m *GreetManyTimesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetService_LongGreet_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).LongGreet(&greetServiceLongGreetServer{stream})
}

type GreetService_LongGreetServer interface {
	SendAndClose(*LongGreetResponse) error
	Recv() (*LongGreetRequest, error)
	grpc.ServerStream
}

type greetServiceLongGreetServer struct {
	grpc.ServerStream
}

func (x *greetServiceLongGreetServer) SendAndClose(m *LongGreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceLongGreetServer) Recv() (*LongGreetRequest, error) {
	m := new(LongGreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetService_BiDiGreet_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).BiDiGreet(&greetServiceBiDiGreetServer{stream})
}

type GreetService_BiDiGreetServer interface {
	Send(*ManyResponse) error
	Recv() (*ManyRequest, error)
	grpc.ServerStream
}

type greetServiceBiDiGreetServer struct {
	grpc.ServerStream
}

func (x *greetServiceBiDiGreetServer) Send(m *ManyResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceBiDiGreetServer) Recv() (*ManyRequest, error) {
	m := new(ManyRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetService_Greet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetManyTimes",
			Handler:       _GreetService_GreetManyTimes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LongGreet",
			Handler:       _GreetService_LongGreet_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BiDiGreet",
			Handler:       _GreetService_BiDiGreet_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "greetpb/greet.proto",
}

func init() { proto.RegisterFile("greetpb/greet.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4d, 0x4b, 0xf3, 0x40,
	0x10, 0xc7, 0x9f, 0x14, 0xda, 0xa7, 0x3b, 0xad, 0x6f, 0x13, 0xad, 0x25, 0x5a, 0x90, 0x1c, 0xb4,
	0x50, 0xa8, 0xa1, 0x7a, 0xaa, 0x07, 0x21, 0x14, 0xbd, 0xa8, 0x60, 0xf4, 0xe4, 0x45, 0x52, 0x19,
	0xc3, 0x42, 0xf3, 0x62, 0x36, 0x15, 0xfc, 0x24, 0x7e, 0x5d, 0xc9, 0x66, 0x53, 0x93, 0x88, 0x16,
	0x72, 0x4a, 0x66, 0xe6, 0x3f, 0xbf, 0xff, 0xec, 0x2c, 0x0b, 0xba, 0x17, 0x13, 0x25, 0xd1, 0xfc,
	0x54, 0x7e, 0xc7, 0x51, 0x1c, 0x26, 0x21, 0x36, 0x65, 0x60, 0x5e, 0x41, 0xfb, 0x3a, 0xfd, 0xe1,
	0x81, 0x87, 0x03, 0x80, 0x57, 0x1e, 0x8b, 0xe4, 0x39, 0x70, 0x7d, 0xea, 0x6b, 0x47, 0xda, 0x90,
	0x39, 0x4c, 0x66, 0xee, 0x5c, 0x9f, 0xf0, 0x00, 0xd8, 0xc2, 0xcd, 0xab, 0x0d, 0x59, 0x6d, 0xa7,
	0x89, 0xb4, 0x68, 0x5e, 0x40, 0x57, 0x72, 0x1c, 0x7a, 0x5b, 0x92, 0x48, 0x70, 0x04, 0x6d, 0x4f,
	0x71, 0x25, 0xa9, 0x33, 0xd9, 0x1a, 0x67, 0xf6, 0xb9, 0x9d, 0xb3, 0x12, 0x98, 0x27, 0xb0, 0xa1,
	0x9a, 0x45, 0x14, 0x06, 0x82, 0xb0, 0x07, 0xad, 0x98, 0xc4, 0x72, 0x91, 0xa8, 0x29, 0x54, 0x64,
	0xce, 0x60, 0x4f, 0x0a, 0x6f, 0xdd, 0xe0, 0xe3, 0x91, 0xfb, 0x24, 0x6a, 0xd9, 0x59, 0xd0, 0xab,
	0x52, 0xd6, 0xf8, 0x5e, 0xc2, 0xf6, 0x4d, 0x18, 0x78, 0xf5, 0x4f, 0x38, 0x82, 0x9d, 0x02, 0x60,
	0x8d, 0xdb, 0x14, 0x3a, 0xe9, 0x68, 0xb5, 0x8c, 0x8e, 0xa1, 0x9b, 0xf5, 0xfe, 0xed, 0x31, 0xf9,
	0x6c, 0xa8, 0x0b, 0x7b, 0xa0, 0xf8, 0x9d, 0xbf, 0x10, 0x9e, 0x43, 0x53, 0xc6, 0xa8, 0x17, 0xe1,
	0x6a, 0x06, 0x63, 0xb7, 0x9c, 0xcc, 0xe0, 0xe6, 0x3f, 0xbc, 0x87, 0xcd, 0xf2, 0x2a, 0xf1, 0xb0,
	0xa8, 0xac, 0xde, 0x93, 0x31, 0xf8, 0xa5, 0x9a, 0x03, 0x2d, 0x0d, 0x6d, 0x60, 0xab, 0x55, 0xe1,
	0xbe, 0xd2, 0x57, 0xb7, 0x6f, 0xf4, 0x7f, 0x16, 0x72, 0xc6, 0x50, 0xc3, 0x29, 0x30, 0x9b, 0xcf,
	0x78, 0xc6, 0x40, 0x25, 0x2d, 0xec, 0xd4, 0xd0, 0x4b, 0xb9, 0xef, 0x4e, 0x4b, 0xb3, 0xd9, 0xd3,
	0x7f, 0xf5, 0x5e, 0xe6, 0x2d, 0xf9, 0x54, 0xce, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xed,
	0xd4, 0x80, 0x41, 0x03, 0x00, 0x00,
}
