// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: recommendation.proto

package Recommendation

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
	Recommendation_GetRecommendationsCards_FullMethodName           = "/Dressy.Recommendation.Recommendation/GetRecommendationsCards"
	Recommendation_UpdateClientEmbeddingStringVector_FullMethodName = "/Dressy.Recommendation.Recommendation/UpdateClientEmbeddingStringVector"
	Recommendation_UpdateClientEmbeddingFloatVector_FullMethodName  = "/Dressy.Recommendation.Recommendation/UpdateClientEmbeddingFloatVector"
)

// RecommendationClient is the client API for Recommendation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecommendationClient interface {
	GetRecommendationsCards(ctx context.Context, in *RecommendationRequest, opts ...grpc.CallOption) (*RecommendationResponse, error)
	UpdateClientEmbeddingStringVector(ctx context.Context, in *ClientEmbeddingStringRequest, opts ...grpc.CallOption) (*ClientEmbeddingResponse, error)
	UpdateClientEmbeddingFloatVector(ctx context.Context, in *ClientEmbeddingFloatRequest, opts ...grpc.CallOption) (*ClientEmbeddingResponse, error)
}

type recommendationClient struct {
	cc grpc.ClientConnInterface
}

func NewRecommendationClient(cc grpc.ClientConnInterface) RecommendationClient {
	return &recommendationClient{cc}
}

func (c *recommendationClient) GetRecommendationsCards(ctx context.Context, in *RecommendationRequest, opts ...grpc.CallOption) (*RecommendationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RecommendationResponse)
	err := c.cc.Invoke(ctx, Recommendation_GetRecommendationsCards_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendationClient) UpdateClientEmbeddingStringVector(ctx context.Context, in *ClientEmbeddingStringRequest, opts ...grpc.CallOption) (*ClientEmbeddingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClientEmbeddingResponse)
	err := c.cc.Invoke(ctx, Recommendation_UpdateClientEmbeddingStringVector_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendationClient) UpdateClientEmbeddingFloatVector(ctx context.Context, in *ClientEmbeddingFloatRequest, opts ...grpc.CallOption) (*ClientEmbeddingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClientEmbeddingResponse)
	err := c.cc.Invoke(ctx, Recommendation_UpdateClientEmbeddingFloatVector_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecommendationServer is the server API for Recommendation service.
// All implementations must embed UnimplementedRecommendationServer
// for forward compatibility.
type RecommendationServer interface {
	GetRecommendationsCards(context.Context, *RecommendationRequest) (*RecommendationResponse, error)
	UpdateClientEmbeddingStringVector(context.Context, *ClientEmbeddingStringRequest) (*ClientEmbeddingResponse, error)
	UpdateClientEmbeddingFloatVector(context.Context, *ClientEmbeddingFloatRequest) (*ClientEmbeddingResponse, error)
	mustEmbedUnimplementedRecommendationServer()
}

// UnimplementedRecommendationServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRecommendationServer struct{}

func (UnimplementedRecommendationServer) GetRecommendationsCards(context.Context, *RecommendationRequest) (*RecommendationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecommendationsCards not implemented")
}
func (UnimplementedRecommendationServer) UpdateClientEmbeddingStringVector(context.Context, *ClientEmbeddingStringRequest) (*ClientEmbeddingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClientEmbeddingStringVector not implemented")
}
func (UnimplementedRecommendationServer) UpdateClientEmbeddingFloatVector(context.Context, *ClientEmbeddingFloatRequest) (*ClientEmbeddingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClientEmbeddingFloatVector not implemented")
}
func (UnimplementedRecommendationServer) mustEmbedUnimplementedRecommendationServer() {}
func (UnimplementedRecommendationServer) testEmbeddedByValue()                        {}

// UnsafeRecommendationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecommendationServer will
// result in compilation errors.
type UnsafeRecommendationServer interface {
	mustEmbedUnimplementedRecommendationServer()
}

func RegisterRecommendationServer(s grpc.ServiceRegistrar, srv RecommendationServer) {
	// If the following call pancis, it indicates UnimplementedRecommendationServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Recommendation_ServiceDesc, srv)
}

func _Recommendation_GetRecommendationsCards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecommendationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServer).GetRecommendationsCards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Recommendation_GetRecommendationsCards_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServer).GetRecommendationsCards(ctx, req.(*RecommendationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Recommendation_UpdateClientEmbeddingStringVector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientEmbeddingStringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServer).UpdateClientEmbeddingStringVector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Recommendation_UpdateClientEmbeddingStringVector_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServer).UpdateClientEmbeddingStringVector(ctx, req.(*ClientEmbeddingStringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Recommendation_UpdateClientEmbeddingFloatVector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientEmbeddingFloatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServer).UpdateClientEmbeddingFloatVector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Recommendation_UpdateClientEmbeddingFloatVector_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServer).UpdateClientEmbeddingFloatVector(ctx, req.(*ClientEmbeddingFloatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Recommendation_ServiceDesc is the grpc.ServiceDesc for Recommendation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Recommendation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Dressy.Recommendation.Recommendation",
	HandlerType: (*RecommendationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRecommendationsCards",
			Handler:    _Recommendation_GetRecommendationsCards_Handler,
		},
		{
			MethodName: "UpdateClientEmbeddingStringVector",
			Handler:    _Recommendation_UpdateClientEmbeddingStringVector_Handler,
		},
		{
			MethodName: "UpdateClientEmbeddingFloatVector",
			Handler:    _Recommendation_UpdateClientEmbeddingFloatVector_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recommendation.proto",
}