// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: demo/v1/demo.proto

package demov1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DemoService_WhatIsDemoInfo_FullMethodName = "/demo.v1.DemoService/WhatIsDemoInfo"
)

// DemoServiceClient is the client API for DemoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DemoServiceClient interface {
	WhatIsDemoInfo(ctx context.Context, in *WhatIsDemoInfoRequest, opts ...grpc.CallOption) (*WhatIsDemoInfoResponse, error)
}

type demoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDemoServiceClient(cc grpc.ClientConnInterface) DemoServiceClient {
	return &demoServiceClient{cc}
}

func (c *demoServiceClient) WhatIsDemoInfo(ctx context.Context, in *WhatIsDemoInfoRequest, opts ...grpc.CallOption) (*WhatIsDemoInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WhatIsDemoInfoResponse)
	err := c.cc.Invoke(ctx, DemoService_WhatIsDemoInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemoServiceServer is the server API for DemoService service.
// All implementations should embed UnimplementedDemoServiceServer
// for forward compatibility.
type DemoServiceServer interface {
	WhatIsDemoInfo(context.Context, *WhatIsDemoInfoRequest) (*WhatIsDemoInfoResponse, error)
}

// UnimplementedDemoServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDemoServiceServer struct{}

func (UnimplementedDemoServiceServer) WhatIsDemoInfo(context.Context, *WhatIsDemoInfoRequest) (*WhatIsDemoInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhatIsDemoInfo not implemented")
}
func (UnimplementedDemoServiceServer) testEmbeddedByValue() {}

// UnsafeDemoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DemoServiceServer will
// result in compilation errors.
type UnsafeDemoServiceServer interface {
	mustEmbedUnimplementedDemoServiceServer()
}

func RegisterDemoServiceServer(s grpc.ServiceRegistrar, srv DemoServiceServer) {
	// If the following call pancis, it indicates UnimplementedDemoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DemoService_ServiceDesc, srv)
}

func _DemoService_WhatIsDemoInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhatIsDemoInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).WhatIsDemoInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DemoService_WhatIsDemoInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).WhatIsDemoInfo(ctx, req.(*WhatIsDemoInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DemoService_ServiceDesc is the grpc.ServiceDesc for DemoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DemoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "demo.v1.DemoService",
	HandlerType: (*DemoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WhatIsDemoInfo",
			Handler:    _DemoService_WhatIsDemoInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo/v1/demo.proto",
}
