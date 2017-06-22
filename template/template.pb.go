// Code generated by protoc-gen-go. DO NOT EDIT.
// source: template/template.proto

/*
Package template is a generated protocol buffer package.

It is generated from these files:
	template/template.proto

It has these top-level messages:
	TemplateRequest
	TemplateResponse
*/
package template

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// The request message containing the user's name.
type TemplateRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *TemplateRequest) Reset()                    { *m = TemplateRequest{} }
func (m *TemplateRequest) String() string            { return proto.CompactTextString(m) }
func (*TemplateRequest) ProtoMessage()               {}
func (*TemplateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TemplateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type TemplateResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *TemplateResponse) Reset()                    { *m = TemplateResponse{} }
func (m *TemplateResponse) String() string            { return proto.CompactTextString(m) }
func (*TemplateResponse) ProtoMessage()               {}
func (*TemplateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TemplateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*TemplateRequest)(nil), "template.TemplateRequest")
	proto.RegisterType((*TemplateResponse)(nil), "template.TemplateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	// Sends a greeting
	SendGet(ctx context.Context, in *TemplateRequest, opts ...grpc.CallOption) (*TemplateResponse, error)
	SendPost(ctx context.Context, in *TemplateRequest, opts ...grpc.CallOption) (*TemplateResponse, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SendGet(ctx context.Context, in *TemplateRequest, opts ...grpc.CallOption) (*TemplateResponse, error) {
	out := new(TemplateResponse)
	err := grpc.Invoke(ctx, "/template.Greeter/SendGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SendPost(ctx context.Context, in *TemplateRequest, opts ...grpc.CallOption) (*TemplateResponse, error) {
	out := new(TemplateResponse)
	err := grpc.Invoke(ctx, "/template.Greeter/SendPost", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	// Sends a greeting
	SendGet(context.Context, *TemplateRequest) (*TemplateResponse, error)
	SendPost(context.Context, *TemplateRequest) (*TemplateResponse, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SendGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SendGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/template.Greeter/SendGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SendGet(ctx, req.(*TemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SendPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SendPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/template.Greeter/SendPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SendPost(ctx, req.(*TemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "template.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendGet",
			Handler:    _Greeter_SendGet_Handler,
		},
		{
			MethodName: "SendPost",
			Handler:    _Greeter_SendPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "template/template.proto",
}

func init() { proto.RegisterFile("template/template.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x49, 0xcd, 0x2d,
	0xc8, 0x49, 0x2c, 0x49, 0xd5, 0x87, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0x60,
	0x7c, 0x29, 0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0xfd, 0xc4, 0x82, 0x4c, 0xfd, 0xc4, 0xbc,
	0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88, 0x3a, 0x25, 0x55, 0x2e, 0xfe, 0x10,
	0xa8, 0xca, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc,
	0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49, 0x87, 0x4b, 0x00, 0xa1, 0xac,
	0xb8, 0x00, 0xa8, 0x3f, 0x55, 0x48, 0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x1d, 0xa6,
	0x14, 0xc6, 0x35, 0x5a, 0xc7, 0xc8, 0xc5, 0xee, 0x5e, 0x94, 0x9a, 0x5a, 0x92, 0x5a, 0x24, 0xe4,
	0xc7, 0xc5, 0x1e, 0x9c, 0x9a, 0x97, 0xe2, 0x9e, 0x5a, 0x22, 0x24, 0xa9, 0x07, 0x77, 0x24, 0x9a,
	0x9d, 0x52, 0x52, 0xd8, 0xa4, 0x20, 0xf6, 0x28, 0xf1, 0x34, 0x5d, 0x7e, 0x32, 0x99, 0x89, 0x4d,
	0x88, 0x45, 0x3f, 0x1d, 0x68, 0x48, 0x30, 0x17, 0x07, 0xc8, 0xbc, 0x80, 0xfc, 0x62, 0xb2, 0x0d,
	0x14, 0x00, 0x1b, 0xc8, 0xa5, 0xc4, 0xaa, 0x5f, 0x00, 0x34, 0xc5, 0x8a, 0x51, 0x2b, 0x89, 0x0d,
	0x1c, 0x18, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x9a, 0xe7, 0xc1, 0x4f, 0x01, 0x00,
	0x00,
}
