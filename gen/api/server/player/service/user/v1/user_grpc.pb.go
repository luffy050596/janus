// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: player/service/user/v1/user.proto

package servicev1

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
	UserService_GetById_FullMethodName          = "/server.player.service.user.v1.UserService/GetById"
	UserService_GetBasicById_FullMethodName     = "/server.player.service.user.v1.UserService/GetBasicById"
	UserService_UserListById_FullMethodName     = "/server.player.service.user.v1.UserService/UserListById"
	UserService_GetBasicListById_FullMethodName = "/server.player.service.user.v1.UserService/GetBasicListById"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// User service
// Open to the server cluster
// Provide HTTP and gRPC interfaces
type UserServiceClient interface {
	// Get player data cache by id
	GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetByIdResponse, error)
	// Get player basic data cache by id
	GetBasicById(ctx context.Context, in *GetBasicByIdRequest, opts ...grpc.CallOption) (*GetBasicByIdResponse, error)
	// Get player data cache list by id list
	UserListById(ctx context.Context, in *UserListByIdRequest, opts ...grpc.CallOption) (*UserListByIdResponse, error)
	// Get player basic data cache list by id list
	GetBasicListById(ctx context.Context, in *GetBasicListByIdRequest, opts ...grpc.CallOption) (*GetBasicListByIdResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetByIdResponse)
	err := c.cc.Invoke(ctx, UserService_GetById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetBasicById(ctx context.Context, in *GetBasicByIdRequest, opts ...grpc.CallOption) (*GetBasicByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBasicByIdResponse)
	err := c.cc.Invoke(ctx, UserService_GetBasicById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserListById(ctx context.Context, in *UserListByIdRequest, opts ...grpc.CallOption) (*UserListByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserListByIdResponse)
	err := c.cc.Invoke(ctx, UserService_UserListById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetBasicListById(ctx context.Context, in *GetBasicListByIdRequest, opts ...grpc.CallOption) (*GetBasicListByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBasicListByIdResponse)
	err := c.cc.Invoke(ctx, UserService_GetBasicListById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
//
// User service
// Open to the server cluster
// Provide HTTP and gRPC interfaces
type UserServiceServer interface {
	// Get player data cache by id
	GetById(context.Context, *GetByIdRequest) (*GetByIdResponse, error)
	// Get player basic data cache by id
	GetBasicById(context.Context, *GetBasicByIdRequest) (*GetBasicByIdResponse, error)
	// Get player data cache list by id list
	UserListById(context.Context, *UserListByIdRequest) (*UserListByIdResponse, error)
	// Get player basic data cache list by id list
	GetBasicListById(context.Context, *GetBasicListByIdRequest) (*GetBasicListByIdResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) GetById(context.Context, *GetByIdRequest) (*GetByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedUserServiceServer) GetBasicById(context.Context, *GetBasicByIdRequest) (*GetBasicByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBasicById not implemented")
}
func (UnimplementedUserServiceServer) UserListById(context.Context, *UserListByIdRequest) (*UserListByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserListById not implemented")
}
func (UnimplementedUserServiceServer) GetBasicListById(context.Context, *GetBasicListByIdRequest) (*GetBasicListByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBasicListById not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetById(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetBasicById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBasicByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetBasicById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetBasicById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetBasicById(ctx, req.(*GetBasicByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserListById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserListByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserListById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserListById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserListById(ctx, req.(*UserListByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetBasicListById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBasicListByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetBasicListById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetBasicListById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetBasicListById(ctx, req.(*GetBasicListByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.player.service.user.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetById",
			Handler:    _UserService_GetById_Handler,
		},
		{
			MethodName: "GetBasicById",
			Handler:    _UserService_GetBasicById_Handler,
		},
		{
			MethodName: "UserListById",
			Handler:    _UserService_UserListById_Handler,
		},
		{
			MethodName: "GetBasicListById",
			Handler:    _UserService_GetBasicListById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "player/service/user/v1/user.proto",
}
