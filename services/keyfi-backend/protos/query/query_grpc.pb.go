// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: protos/query/query.proto

package query

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

const (
	QueryService_GetValues_FullMethodName = "/keyfi_protos.query.QueryService/GetValues"
)

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryServiceClient interface {
	GetValues(ctx context.Context, in *GetValuesRequest, opts ...grpc.CallOption) (*GetValuesResponse, error)
}

type queryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryServiceClient(cc grpc.ClientConnInterface) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) GetValues(ctx context.Context, in *GetValuesRequest, opts ...grpc.CallOption) (*GetValuesResponse, error) {
	out := new(GetValuesResponse)
	err := c.cc.Invoke(ctx, QueryService_GetValues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
// All implementations must embed UnimplementedQueryServiceServer
// for forward compatibility
type QueryServiceServer interface {
	GetValues(context.Context, *GetValuesRequest) (*GetValuesResponse, error)
	mustEmbedUnimplementedQueryServiceServer()
}

// UnimplementedQueryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (UnimplementedQueryServiceServer) GetValues(context.Context, *GetValuesRequest) (*GetValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValues not implemented")
}
func (UnimplementedQueryServiceServer) mustEmbedUnimplementedQueryServiceServer() {}

// UnsafeQueryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServiceServer will
// result in compilation errors.
type UnsafeQueryServiceServer interface {
	mustEmbedUnimplementedQueryServiceServer()
}

func RegisterQueryServiceServer(s grpc.ServiceRegistrar, srv QueryServiceServer) {
	s.RegisterService(&QueryService_ServiceDesc, srv)
}

func _QueryService_GetValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).GetValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueryService_GetValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).GetValues(ctx, req.(*GetValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QueryService_ServiceDesc is the grpc.ServiceDesc for QueryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QueryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "keyfi_protos.query.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetValues",
			Handler:    _QueryService_GetValues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/query/query.proto",
}