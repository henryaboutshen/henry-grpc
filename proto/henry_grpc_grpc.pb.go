// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: proto/henry_grpc.proto

package henry_grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type helloClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloClient(cc grpc.ClientConnInterface) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/henry_grpc.Hello/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServer is the server API for Hello service.
// All implementations must embed UnimplementedHelloServer
// for forward compatibility
type HelloServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedHelloServer()
}

// UnimplementedHelloServer must be embedded to have forward compatible implementations.
type UnimplementedHelloServer struct {
}

func (UnimplementedHelloServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedHelloServer) mustEmbedUnimplementedHelloServer() {}

// UnsafeHelloServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServer will
// result in compilation errors.
type UnsafeHelloServer interface {
	mustEmbedUnimplementedHelloServer()
}

func RegisterHelloServer(s grpc.ServiceRegistrar, srv HelloServer) {
	s.RegisterService(&Hello_ServiceDesc, srv)
}

func _Hello_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/henry_grpc.Hello/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Hello_ServiceDesc is the grpc.ServiceDesc for Hello service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hello_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "henry_grpc.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Hello_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/henry_grpc.proto",
}

// HiClient is the client API for Hi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HiClient interface {
	SayHi(ctx context.Context, in *HiRequest, opts ...grpc.CallOption) (*HiReply, error)
}

type hiClient struct {
	cc grpc.ClientConnInterface
}

func NewHiClient(cc grpc.ClientConnInterface) HiClient {
	return &hiClient{cc}
}

func (c *hiClient) SayHi(ctx context.Context, in *HiRequest, opts ...grpc.CallOption) (*HiReply, error) {
	out := new(HiReply)
	err := c.cc.Invoke(ctx, "/henry_grpc.Hi/SayHi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HiServer is the server API for Hi service.
// All implementations must embed UnimplementedHiServer
// for forward compatibility
type HiServer interface {
	SayHi(context.Context, *HiRequest) (*HiReply, error)
	mustEmbedUnimplementedHiServer()
}

// UnimplementedHiServer must be embedded to have forward compatible implementations.
type UnimplementedHiServer struct {
}

func (UnimplementedHiServer) SayHi(context.Context, *HiRequest) (*HiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHi not implemented")
}
func (UnimplementedHiServer) mustEmbedUnimplementedHiServer() {}

// UnsafeHiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HiServer will
// result in compilation errors.
type UnsafeHiServer interface {
	mustEmbedUnimplementedHiServer()
}

func RegisterHiServer(s grpc.ServiceRegistrar, srv HiServer) {
	s.RegisterService(&Hi_ServiceDesc, srv)
}

func _Hi_SayHi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HiServer).SayHi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/henry_grpc.Hi/SayHi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HiServer).SayHi(ctx, req.(*HiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Hi_ServiceDesc is the grpc.ServiceDesc for Hi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "henry_grpc.Hi",
	HandlerType: (*HiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHi",
			Handler:    _Hi_SayHi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/henry_grpc.proto",
}