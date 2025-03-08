// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: chat/v1/chat.proto

package chatv1

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
	ChatService_GenerateMessage_FullMethodName = "/chat.v1.ChatService/GenerateMessage"
	ChatService_StreamMessage_FullMethodName   = "/chat.v1.ChatService/StreamMessage"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	GenerateMessage(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatResponse, error)
	StreamMessage(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ChatResponse], error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) GenerateMessage(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChatResponse)
	err := c.cc.Invoke(ctx, ChatService_GenerateMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) StreamMessage(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ChatResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], ChatService_StreamMessage_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ChatRequest, ChatResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatService_StreamMessageClient = grpc.ServerStreamingClient[ChatResponse]

// ChatServiceServer is the server API for ChatService service.
// All implementations should embed UnimplementedChatServiceServer
// for forward compatibility.
type ChatServiceServer interface {
	GenerateMessage(context.Context, *ChatRequest) (*ChatResponse, error)
	StreamMessage(*ChatRequest, grpc.ServerStreamingServer[ChatResponse]) error
}

// UnimplementedChatServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChatServiceServer struct{}

func (UnimplementedChatServiceServer) GenerateMessage(context.Context, *ChatRequest) (*ChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateMessage not implemented")
}
func (UnimplementedChatServiceServer) StreamMessage(*ChatRequest, grpc.ServerStreamingServer[ChatResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamMessage not implemented")
}
func (UnimplementedChatServiceServer) testEmbeddedByValue() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	// If the following call pancis, it indicates UnimplementedChatServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_GenerateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GenerateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GenerateMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GenerateMessage(ctx, req.(*ChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_StreamMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ChatRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).StreamMessage(m, &grpc.GenericServerStream[ChatRequest, ChatResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatService_StreamMessageServer = grpc.ServerStreamingServer[ChatResponse]

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.v1.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateMessage",
			Handler:    _ChatService_GenerateMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMessage",
			Handler:       _ChatService_StreamMessage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat/v1/chat.proto",
}
