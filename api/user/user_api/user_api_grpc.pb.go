// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: user/user_api.proto

package user_api

import (
	context "context"
	response_model "github.com/Alexander272/sealur_proto/api/pro/models/response_model"
	user_model "github.com/Alexander272/sealur_proto/api/user/models/user_model"
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
	Get(ctx context.Context, in *GetUser, opts ...grpc.CallOption) (*user_model.User, error)
	GetFull(ctx context.Context, in *GetUser, opts ...grpc.CallOption) (*user_model.FullUser, error)
	GetByEmail(ctx context.Context, in *GetUserByEmail, opts ...grpc.CallOption) (*user_model.User, error)
	GetByParam(ctx context.Context, in *GetUsersByParam, opts ...grpc.CallOption) (*AnalyticsUsers, error)
	GetAll(ctx context.Context, in *GetAllUser, opts ...grpc.CallOption) (*Users, error)
	GetNew(ctx context.Context, in *GetNewUser, opts ...grpc.CallOption) (*Users, error)
	GetManagers(ctx context.Context, in *GetNewUser, opts ...grpc.CallOption) (*Users, error)
	GetManagerEmail(ctx context.Context, in *GetUser, opts ...grpc.CallOption) (*Manager, error)
	GetAnalytics(ctx context.Context, in *GetUserAnalytics, opts ...grpc.CallOption) (*Analytics, error)
	Create(ctx context.Context, in *CreateUser, opts ...grpc.CallOption) (*response_model.IdResponse, error)
	Confirm(ctx context.Context, in *ConfirmUser, opts ...grpc.CallOption) (*user_model.User, error)
	Update(ctx context.Context, in *UpdateUser, opts ...grpc.CallOption) (*response_model.Response, error)
	SetManager(ctx context.Context, in *UserManager, opts ...grpc.CallOption) (*response_model.Response, error)
	Delete(ctx context.Context, in *DeleteUser, opts ...grpc.CallOption) (*response_model.Response, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Get(ctx context.Context, in *GetUser, opts ...grpc.CallOption) (*user_model.User, error) {
	out := new(user_model.User)
	err := c.cc.Invoke(ctx, "/user_api.UserService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetFull(ctx context.Context, in *GetUser, opts ...grpc.CallOption) (*user_model.FullUser, error) {
	out := new(user_model.FullUser)
	err := c.cc.Invoke(ctx, "/user_api.UserService/GetFull", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetByEmail(ctx context.Context, in *GetUserByEmail, opts ...grpc.CallOption) (*user_model.User, error) {
	out := new(user_model.User)
	err := c.cc.Invoke(ctx, "/user_api.UserService/GetByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetByParam(ctx context.Context, in *GetUsersByParam, opts ...grpc.CallOption) (*AnalyticsUsers, error) {
	out := new(AnalyticsUsers)
	err := c.cc.Invoke(ctx, "/user_api.UserService/GetByParam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAll(ctx context.Context, in *GetAllUser, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/user_api.UserService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetNew(ctx context.Context, in *GetNewUser, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/user_api.UserService/GetNew", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetManagers(ctx context.Context, in *GetNewUser, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/user_api.UserService/GetManagers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetManagerEmail(ctx context.Context, in *GetUser, opts ...grpc.CallOption) (*Manager, error) {
	out := new(Manager)
	err := c.cc.Invoke(ctx, "/user_api.UserService/GetManagerEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAnalytics(ctx context.Context, in *GetUserAnalytics, opts ...grpc.CallOption) (*Analytics, error) {
	out := new(Analytics)
	err := c.cc.Invoke(ctx, "/user_api.UserService/GetAnalytics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Create(ctx context.Context, in *CreateUser, opts ...grpc.CallOption) (*response_model.IdResponse, error) {
	out := new(response_model.IdResponse)
	err := c.cc.Invoke(ctx, "/user_api.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Confirm(ctx context.Context, in *ConfirmUser, opts ...grpc.CallOption) (*user_model.User, error) {
	out := new(user_model.User)
	err := c.cc.Invoke(ctx, "/user_api.UserService/Confirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UpdateUser, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/user_api.UserService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SetManager(ctx context.Context, in *UserManager, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/user_api.UserService/SetManager", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *DeleteUser, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/user_api.UserService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Get(context.Context, *GetUser) (*user_model.User, error)
	GetFull(context.Context, *GetUser) (*user_model.FullUser, error)
	GetByEmail(context.Context, *GetUserByEmail) (*user_model.User, error)
	GetByParam(context.Context, *GetUsersByParam) (*AnalyticsUsers, error)
	GetAll(context.Context, *GetAllUser) (*Users, error)
	GetNew(context.Context, *GetNewUser) (*Users, error)
	GetManagers(context.Context, *GetNewUser) (*Users, error)
	GetManagerEmail(context.Context, *GetUser) (*Manager, error)
	GetAnalytics(context.Context, *GetUserAnalytics) (*Analytics, error)
	Create(context.Context, *CreateUser) (*response_model.IdResponse, error)
	Confirm(context.Context, *ConfirmUser) (*user_model.User, error)
	Update(context.Context, *UpdateUser) (*response_model.Response, error)
	SetManager(context.Context, *UserManager) (*response_model.Response, error)
	Delete(context.Context, *DeleteUser) (*response_model.Response, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Get(context.Context, *GetUser) (*user_model.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUserServiceServer) GetFull(context.Context, *GetUser) (*user_model.FullUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFull not implemented")
}
func (UnimplementedUserServiceServer) GetByEmail(context.Context, *GetUserByEmail) (*user_model.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByEmail not implemented")
}
func (UnimplementedUserServiceServer) GetByParam(context.Context, *GetUsersByParam) (*AnalyticsUsers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByParam not implemented")
}
func (UnimplementedUserServiceServer) GetAll(context.Context, *GetAllUser) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedUserServiceServer) GetNew(context.Context, *GetNewUser) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNew not implemented")
}
func (UnimplementedUserServiceServer) GetManagers(context.Context, *GetNewUser) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManagers not implemented")
}
func (UnimplementedUserServiceServer) GetManagerEmail(context.Context, *GetUser) (*Manager, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManagerEmail not implemented")
}
func (UnimplementedUserServiceServer) GetAnalytics(context.Context, *GetUserAnalytics) (*Analytics, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnalytics not implemented")
}
func (UnimplementedUserServiceServer) Create(context.Context, *CreateUser) (*response_model.IdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserServiceServer) Confirm(context.Context, *ConfirmUser) (*user_model.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Confirm not implemented")
}
func (UnimplementedUserServiceServer) Update(context.Context, *UpdateUser) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserServiceServer) SetManager(context.Context, *UserManager) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetManager not implemented")
}
func (UnimplementedUserServiceServer) Delete(context.Context, *DeleteUser) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
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

func _UserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api.UserService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Get(ctx, req.(*GetUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetFull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetFull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api.UserService/GetFull",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetFull(ctx, req.(*GetUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByEmail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api.UserService/GetByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetByEmail(ctx, req.(*GetUserByEmail))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetByParam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersByParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetByParam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api.UserService/GetByParam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetByParam(ctx, req.(*GetUsersByParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api.UserService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAll(ctx, req.(*GetAllUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetNew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetNew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api.UserService/GetNew",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetNew(ctx, req.(*GetNewUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetManagers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetManagers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api.UserService/GetManagers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetManagers(ctx, req.(*GetNewUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetManagerEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetManagerEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api.UserService/GetManagerEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetManagerEmail(ctx, req.(*GetUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAnalytics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAnalytics)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAnalytics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api.UserService/GetAnalytics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAnalytics(ctx, req.(*GetUserAnalytics))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*CreateUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Confirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Confirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api.UserService/Confirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Confirm(ctx, req.(*ConfirmUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api.UserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UpdateUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SetManager_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserManager)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SetManager(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api.UserService/SetManager",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SetManager(ctx, req.(*UserManager))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_api.UserService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*DeleteUser))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_api.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _UserService_Get_Handler,
		},
		{
			MethodName: "GetFull",
			Handler:    _UserService_GetFull_Handler,
		},
		{
			MethodName: "GetByEmail",
			Handler:    _UserService_GetByEmail_Handler,
		},
		{
			MethodName: "GetByParam",
			Handler:    _UserService_GetByParam_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _UserService_GetAll_Handler,
		},
		{
			MethodName: "GetNew",
			Handler:    _UserService_GetNew_Handler,
		},
		{
			MethodName: "GetManagers",
			Handler:    _UserService_GetManagers_Handler,
		},
		{
			MethodName: "GetManagerEmail",
			Handler:    _UserService_GetManagerEmail_Handler,
		},
		{
			MethodName: "GetAnalytics",
			Handler:    _UserService_GetAnalytics_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Confirm",
			Handler:    _UserService_Confirm_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "SetManager",
			Handler:    _UserService_SetManager_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user_api.proto",
}
