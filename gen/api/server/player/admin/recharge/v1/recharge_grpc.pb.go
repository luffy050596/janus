// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: player/admin/recharge/v1/recharge.proto

package adminv1

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
	RechargeAdmin_GetOrderList_FullMethodName            = "/server.player.admin.recharge.v1.RechargeAdmin/GetOrderList"
	RechargeAdmin_GetOrderById_FullMethodName            = "/server.player.admin.recharge.v1.RechargeAdmin/GetOrderById"
	RechargeAdmin_UpdateOrderAckStateById_FullMethodName = "/server.player.admin.recharge.v1.RechargeAdmin/UpdateOrderAckStateById"
)

// RechargeAdminClient is the client API for RechargeAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Recharge admin service
// Open to the server cluster
// Provide HTTP and gRPC interfaces
type RechargeAdminClient interface {
	// Query order list by page
	GetOrderList(ctx context.Context, in *GetOrderListRequest, opts ...grpc.CallOption) (*GetOrderListResponse, error)
	// Query order by id
	GetOrderById(ctx context.Context, in *GetOrderByIdRequest, opts ...grpc.CallOption) (*GetOrderByIdResponse, error)
	// Update order confirmation status by id
	UpdateOrderAckStateById(ctx context.Context, in *UpdateOrderAckStateByIdRequest, opts ...grpc.CallOption) (*UpdateOrderAckStateByIdResponse, error)
}

type rechargeAdminClient struct {
	cc grpc.ClientConnInterface
}

func NewRechargeAdminClient(cc grpc.ClientConnInterface) RechargeAdminClient {
	return &rechargeAdminClient{cc}
}

func (c *rechargeAdminClient) GetOrderList(ctx context.Context, in *GetOrderListRequest, opts ...grpc.CallOption) (*GetOrderListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOrderListResponse)
	err := c.cc.Invoke(ctx, RechargeAdmin_GetOrderList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rechargeAdminClient) GetOrderById(ctx context.Context, in *GetOrderByIdRequest, opts ...grpc.CallOption) (*GetOrderByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOrderByIdResponse)
	err := c.cc.Invoke(ctx, RechargeAdmin_GetOrderById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rechargeAdminClient) UpdateOrderAckStateById(ctx context.Context, in *UpdateOrderAckStateByIdRequest, opts ...grpc.CallOption) (*UpdateOrderAckStateByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateOrderAckStateByIdResponse)
	err := c.cc.Invoke(ctx, RechargeAdmin_UpdateOrderAckStateById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RechargeAdminServer is the server API for RechargeAdmin service.
// All implementations must embed UnimplementedRechargeAdminServer
// for forward compatibility.
//
// Recharge admin service
// Open to the server cluster
// Provide HTTP and gRPC interfaces
type RechargeAdminServer interface {
	// Query order list by page
	GetOrderList(context.Context, *GetOrderListRequest) (*GetOrderListResponse, error)
	// Query order by id
	GetOrderById(context.Context, *GetOrderByIdRequest) (*GetOrderByIdResponse, error)
	// Update order confirmation status by id
	UpdateOrderAckStateById(context.Context, *UpdateOrderAckStateByIdRequest) (*UpdateOrderAckStateByIdResponse, error)
	mustEmbedUnimplementedRechargeAdminServer()
}

// UnimplementedRechargeAdminServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRechargeAdminServer struct{}

func (UnimplementedRechargeAdminServer) GetOrderList(context.Context, *GetOrderListRequest) (*GetOrderListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderList not implemented")
}
func (UnimplementedRechargeAdminServer) GetOrderById(context.Context, *GetOrderByIdRequest) (*GetOrderByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderById not implemented")
}
func (UnimplementedRechargeAdminServer) UpdateOrderAckStateById(context.Context, *UpdateOrderAckStateByIdRequest) (*UpdateOrderAckStateByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderAckStateById not implemented")
}
func (UnimplementedRechargeAdminServer) mustEmbedUnimplementedRechargeAdminServer() {}
func (UnimplementedRechargeAdminServer) testEmbeddedByValue()                       {}

// UnsafeRechargeAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RechargeAdminServer will
// result in compilation errors.
type UnsafeRechargeAdminServer interface {
	mustEmbedUnimplementedRechargeAdminServer()
}

func RegisterRechargeAdminServer(s grpc.ServiceRegistrar, srv RechargeAdminServer) {
	// If the following call pancis, it indicates UnimplementedRechargeAdminServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RechargeAdmin_ServiceDesc, srv)
}

func _RechargeAdmin_GetOrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RechargeAdminServer).GetOrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RechargeAdmin_GetOrderList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RechargeAdminServer).GetOrderList(ctx, req.(*GetOrderListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RechargeAdmin_GetOrderById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RechargeAdminServer).GetOrderById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RechargeAdmin_GetOrderById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RechargeAdminServer).GetOrderById(ctx, req.(*GetOrderByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RechargeAdmin_UpdateOrderAckStateById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderAckStateByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RechargeAdminServer).UpdateOrderAckStateById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RechargeAdmin_UpdateOrderAckStateById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RechargeAdminServer).UpdateOrderAckStateById(ctx, req.(*UpdateOrderAckStateByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RechargeAdmin_ServiceDesc is the grpc.ServiceDesc for RechargeAdmin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RechargeAdmin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.player.admin.recharge.v1.RechargeAdmin",
	HandlerType: (*RechargeAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrderList",
			Handler:    _RechargeAdmin_GetOrderList_Handler,
		},
		{
			MethodName: "GetOrderById",
			Handler:    _RechargeAdmin_GetOrderById_Handler,
		},
		{
			MethodName: "UpdateOrderAckStateById",
			Handler:    _RechargeAdmin_UpdateOrderAckStateById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "player/admin/recharge/v1/recharge.proto",
}
