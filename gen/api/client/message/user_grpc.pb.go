// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: message/user.proto

package climsg

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
	UserTCPService_Login_FullMethodName      = "/message.UserTCPService/Login"
	UserTCPService_UpdateName_FullMethodName = "/message.UserTCPService/UpdateName"
	UserTCPService_SetGender_FullMethodName  = "/message.UserTCPService/SetGender"
)

// UserTCPServiceClient is the client API for UserTCPService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// ---------- Service ----------
// User service
type UserTCPServiceClient interface {
	// Login
	Login(ctx context.Context, in *CSLogin, opts ...grpc.CallOption) (*SCLogin, error)
	// Update name
	UpdateName(ctx context.Context, in *CSUpdateName, opts ...grpc.CallOption) (*SCUpdateName, error)
	// Set gender
	SetGender(ctx context.Context, in *CSSetGender, opts ...grpc.CallOption) (*SCSetGender, error)
}

type userTCPServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserTCPServiceClient(cc grpc.ClientConnInterface) UserTCPServiceClient {
	return &userTCPServiceClient{cc}
}

func (c *userTCPServiceClient) Login(ctx context.Context, in *CSLogin, opts ...grpc.CallOption) (*SCLogin, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SCLogin)
	err := c.cc.Invoke(ctx, UserTCPService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userTCPServiceClient) UpdateName(ctx context.Context, in *CSUpdateName, opts ...grpc.CallOption) (*SCUpdateName, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SCUpdateName)
	err := c.cc.Invoke(ctx, UserTCPService_UpdateName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userTCPServiceClient) SetGender(ctx context.Context, in *CSSetGender, opts ...grpc.CallOption) (*SCSetGender, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SCSetGender)
	err := c.cc.Invoke(ctx, UserTCPService_SetGender_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserTCPServiceServer is the server API for UserTCPService service.
// All implementations must embed UnimplementedUserTCPServiceServer
// for forward compatibility.
//
// ---------- Service ----------
// User service
type UserTCPServiceServer interface {
	// Login
	Login(context.Context, *CSLogin) (*SCLogin, error)
	// Update name
	UpdateName(context.Context, *CSUpdateName) (*SCUpdateName, error)
	// Set gender
	SetGender(context.Context, *CSSetGender) (*SCSetGender, error)
	mustEmbedUnimplementedUserTCPServiceServer()
}

// UnimplementedUserTCPServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserTCPServiceServer struct{}

func (UnimplementedUserTCPServiceServer) Login(context.Context, *CSLogin) (*SCLogin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserTCPServiceServer) UpdateName(context.Context, *CSUpdateName) (*SCUpdateName, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateName not implemented")
}
func (UnimplementedUserTCPServiceServer) SetGender(context.Context, *CSSetGender) (*SCSetGender, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGender not implemented")
}
func (UnimplementedUserTCPServiceServer) mustEmbedUnimplementedUserTCPServiceServer() {}
func (UnimplementedUserTCPServiceServer) testEmbeddedByValue()                        {}

// UnsafeUserTCPServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserTCPServiceServer will
// result in compilation errors.
type UnsafeUserTCPServiceServer interface {
	mustEmbedUnimplementedUserTCPServiceServer()
}

func RegisterUserTCPServiceServer(s grpc.ServiceRegistrar, srv UserTCPServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserTCPServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserTCPService_ServiceDesc, srv)
}

func _UserTCPService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserTCPServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserTCPService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserTCPServiceServer).Login(ctx, req.(*CSLogin))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserTCPService_UpdateName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSUpdateName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserTCPServiceServer).UpdateName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserTCPService_UpdateName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserTCPServiceServer).UpdateName(ctx, req.(*CSUpdateName))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserTCPService_SetGender_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSSetGender)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserTCPServiceServer).SetGender(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserTCPService_SetGender_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserTCPServiceServer).SetGender(ctx, req.(*CSSetGender))
	}
	return interceptor(ctx, in, info, handler)
}

// UserTCPService_ServiceDesc is the grpc.ServiceDesc for UserTCPService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserTCPService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.UserTCPService",
	HandlerType: (*UserTCPServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserTCPService_Login_Handler,
		},
		{
			MethodName: "UpdateName",
			Handler:    _UserTCPService_UpdateName_Handler,
		},
		{
			MethodName: "SetGender",
			Handler:    _UserTCPService_SetGender_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message/user.proto",
}
