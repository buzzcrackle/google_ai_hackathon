// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: protos/ai/ai.proto

package ai

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

// AIServiceClient is the client API for AIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AIServiceClient interface {
	SinglePrompt(ctx context.Context, in *SinglePromptRequest, opts ...grpc.CallOption) (*SinglePromptResponse, error)
}

type aIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAIServiceClient(cc grpc.ClientConnInterface) AIServiceClient {
	return &aIServiceClient{cc}
}

func (c *aIServiceClient) SinglePrompt(ctx context.Context, in *SinglePromptRequest, opts ...grpc.CallOption) (*SinglePromptResponse, error) {
	out := new(SinglePromptResponse)
	err := c.cc.Invoke(ctx, "/keyfi_protos.ai.AIService/SinglePrompt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AIServiceServer is the server API for AIService service.
// All implementations must embed UnimplementedAIServiceServer
// for forward compatibility
type AIServiceServer interface {
	SinglePrompt(context.Context, *SinglePromptRequest) (*SinglePromptResponse, error)
	mustEmbedUnimplementedAIServiceServer()
}

// UnimplementedAIServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAIServiceServer struct {
}

func (UnimplementedAIServiceServer) SinglePrompt(context.Context, *SinglePromptRequest) (*SinglePromptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SinglePrompt not implemented")
}
func (UnimplementedAIServiceServer) mustEmbedUnimplementedAIServiceServer() {}

// UnsafeAIServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AIServiceServer will
// result in compilation errors.
type UnsafeAIServiceServer interface {
	mustEmbedUnimplementedAIServiceServer()
}

func RegisterAIServiceServer(s grpc.ServiceRegistrar, srv AIServiceServer) {
	s.RegisterService(&AIService_ServiceDesc, srv)
}

func _AIService_SinglePrompt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SinglePromptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIServiceServer).SinglePrompt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyfi_protos.ai.AIService/SinglePrompt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIServiceServer).SinglePrompt(ctx, req.(*SinglePromptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AIService_ServiceDesc is the grpc.ServiceDesc for AIService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AIService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "keyfi_protos.ai.AIService",
	HandlerType: (*AIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SinglePrompt",
			Handler:    _AIService_SinglePrompt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/ai/ai.proto",
}
