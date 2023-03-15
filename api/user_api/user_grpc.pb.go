// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: user_api/user.proto

package user_api_old

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// User Service ----------------------------------------------------------------------
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetAllUsers(ctx context.Context, in *GetAllUserRequest, opts ...grpc.CallOption) (*UsersResponse, error)
	GetNewUsers(ctx context.Context, in *GetNewUserRequest, opts ...grpc.CallOption) (*UsersResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	ConfirmUser(ctx context.Context, in *ConfirmUserRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*IdResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*IdResponse, error)
	RejectUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*IdResponse, error)
	// Role Service ----------------------------------------------------------------------
	GetRoles(ctx context.Context, in *GetRolesRequest, opts ...grpc.CallOption) (*RolesResponse, error)
	CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*IdResponse, error)
	DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*IdResponse, error)
	// IP service ------------------------------------------------------------
	AddIp(ctx context.Context, in *AddIpRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/user_api_old.UserService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user_api_old.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllUsers(ctx context.Context, in *GetAllUserRequest, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, "/user_api_old.UserService/GetAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetNewUsers(ctx context.Context, in *GetNewUserRequest, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, "/user_api_old.UserService/GetNewUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/user_api_old.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ConfirmUser(ctx context.Context, in *ConfirmUserRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/user_api_old.UserService/ConfirmUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*IdResponse, error) {
	out := new(IdResponse)
	err := c.cc.Invoke(ctx, "/user_api_old.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*IdResponse, error) {
	out := new(IdResponse)
	err := c.cc.Invoke(ctx, "/user_api_old.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RejectUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*IdResponse, error) {
	out := new(IdResponse)
	err := c.cc.Invoke(ctx, "/user_api_old.UserService/RejectUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetRoles(ctx context.Context, in *GetRolesRequest, opts ...grpc.CallOption) (*RolesResponse, error) {
	out := new(RolesResponse)
	err := c.cc.Invoke(ctx, "/user_api_old.UserService/GetRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/user_api_old.UserService/CreateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*IdResponse, error) {
	out := new(IdResponse)
	err := c.cc.Invoke(ctx, "/user_api_old.UserService/UpdateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*IdResponse, error) {
	out := new(IdResponse)
	err := c.cc.Invoke(ctx, "/user_api_old.UserService/DeleteRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddIp(ctx context.Context, in *AddIpRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/user_api_old.UserService/AddIp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// User Service ----------------------------------------------------------------------
	GetUser(context.Context, *GetUserRequest) (*UserResponse, error)
	GetAllUsers(context.Context, *GetAllUserRequest) (*UsersResponse, error)
	GetNewUsers(context.Context, *GetNewUserRequest) (*UsersResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*SuccessResponse, error)
	ConfirmUser(context.Context, *ConfirmUserRequest) (*SuccessResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*IdResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*IdResponse, error)
	RejectUser(context.Context, *DeleteUserRequest) (*IdResponse, error)
	// Role Service ----------------------------------------------------------------------
	GetRoles(context.Context, *GetRolesRequest) (*RolesResponse, error)
	CreateRole(context.Context, *CreateRoleRequest) (*SuccessResponse, error)
	UpdateRole(context.Context, *UpdateRoleRequest) (*IdResponse, error)
	DeleteRole(context.Context, *DeleteRoleRequest) (*IdResponse, error)
	// IP service ------------------------------------------------------------
	AddIp(context.Context, *AddIpRequest) (*SuccessResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) GetAllUsers(context.Context, *GetAllUserRequest) (*UsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedUserServiceServer) GetNewUsers(context.Context, *GetNewUserRequest) (*UsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewUsers not implemented")
}
func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) ConfirmUser(context.Context, *ConfirmUserRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*IdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*IdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) RejectUser(context.Context, *DeleteUserRequest) (*IdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectUser not implemented")
}
func (UnimplementedUserServiceServer) GetRoles(context.Context, *GetRolesRequest) (*RolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoles not implemented")
}
func (UnimplementedUserServiceServer) CreateRole(context.Context, *CreateRoleRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedUserServiceServer) UpdateRole(context.Context, *UpdateRoleRequest) (*IdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedUserServiceServer) DeleteRole(context.Context, *DeleteRoleRequest) (*IdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}
func (UnimplementedUserServiceServer) AddIp(context.Context, *AddIpRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddIp not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api_old.UserService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api_old.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api_old.UserService/GetAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllUsers(ctx, req.(*GetAllUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetNewUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetNewUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api_old.UserService/GetNewUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetNewUsers(ctx, req.(*GetNewUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api_old.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ConfirmUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ConfirmUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api_old.UserService/ConfirmUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ConfirmUser(ctx, req.(*ConfirmUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api_old.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api_old.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RejectUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RejectUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api_old.UserService/RejectUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RejectUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api_old.UserService/GetRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetRoles(ctx, req.(*GetRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api_old.UserService/CreateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateRole(ctx, req.(*CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api_old.UserService/UpdateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateRole(ctx, req.(*UpdateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api_old.UserService/DeleteRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteRole(ctx, req.(*DeleteRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddIp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddIpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddIp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api_old.UserService/AddIp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddIp(ctx, req.(*AddIpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_api_old.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _UserService_Ping_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _UserService_GetAllUsers_Handler,
		},
		{
			MethodName: "GetNewUsers",
			Handler:    _UserService_GetNewUsers_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "ConfirmUser",
			Handler:    _UserService_ConfirmUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
		{
			MethodName: "RejectUser",
			Handler:    _UserService_RejectUser_Handler,
		},
		{
			MethodName: "GetRoles",
			Handler:    _UserService_GetRoles_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _UserService_CreateRole_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _UserService_UpdateRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _UserService_DeleteRole_Handler,
		},
		{
			MethodName: "AddIp",
			Handler:    _UserService_AddIp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_api/user.proto",
}
