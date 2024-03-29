// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: pro/flange_type_api.proto

package flange_type_api

import (
	context "context"
	response_model "github.com/Alexander272/sealur_proto/api/pro/models/response_model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	FlangeTypeService_Get_FullMethodName           = "/flange_type_api.FlangeTypeService/Get"
	FlangeTypeService_Create_FullMethodName        = "/flange_type_api.FlangeTypeService/Create"
	FlangeTypeService_CreateSeveral_FullMethodName = "/flange_type_api.FlangeTypeService/CreateSeveral"
	FlangeTypeService_Update_FullMethodName        = "/flange_type_api.FlangeTypeService/Update"
	FlangeTypeService_Delete_FullMethodName        = "/flange_type_api.FlangeTypeService/Delete"
)

// FlangeTypeServiceClient is the client API for FlangeTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlangeTypeServiceClient interface {
	Get(ctx context.Context, in *GetFlangeType, opts ...grpc.CallOption) (*FlangeType, error)
	Create(ctx context.Context, in *CreateFlangeType, opts ...grpc.CallOption) (*response_model.Response, error)
	CreateSeveral(ctx context.Context, in *CreateSeveralFlangeType, opts ...grpc.CallOption) (*response_model.Response, error)
	Update(ctx context.Context, in *UpdateFlangeType, opts ...grpc.CallOption) (*response_model.Response, error)
	Delete(ctx context.Context, in *DeleteFlangeType, opts ...grpc.CallOption) (*response_model.Response, error)
}

type flangeTypeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFlangeTypeServiceClient(cc grpc.ClientConnInterface) FlangeTypeServiceClient {
	return &flangeTypeServiceClient{cc}
}

func (c *flangeTypeServiceClient) Get(ctx context.Context, in *GetFlangeType, opts ...grpc.CallOption) (*FlangeType, error) {
	out := new(FlangeType)
	err := c.cc.Invoke(ctx, FlangeTypeService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flangeTypeServiceClient) Create(ctx context.Context, in *CreateFlangeType, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, FlangeTypeService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flangeTypeServiceClient) CreateSeveral(ctx context.Context, in *CreateSeveralFlangeType, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, FlangeTypeService_CreateSeveral_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flangeTypeServiceClient) Update(ctx context.Context, in *UpdateFlangeType, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, FlangeTypeService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flangeTypeServiceClient) Delete(ctx context.Context, in *DeleteFlangeType, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, FlangeTypeService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlangeTypeServiceServer is the server API for FlangeTypeService service.
// All implementations must embed UnimplementedFlangeTypeServiceServer
// for forward compatibility
type FlangeTypeServiceServer interface {
	Get(context.Context, *GetFlangeType) (*FlangeType, error)
	Create(context.Context, *CreateFlangeType) (*response_model.Response, error)
	CreateSeveral(context.Context, *CreateSeveralFlangeType) (*response_model.Response, error)
	Update(context.Context, *UpdateFlangeType) (*response_model.Response, error)
	Delete(context.Context, *DeleteFlangeType) (*response_model.Response, error)
	mustEmbedUnimplementedFlangeTypeServiceServer()
}

// UnimplementedFlangeTypeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFlangeTypeServiceServer struct {
}

func (UnimplementedFlangeTypeServiceServer) Get(context.Context, *GetFlangeType) (*FlangeType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedFlangeTypeServiceServer) Create(context.Context, *CreateFlangeType) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedFlangeTypeServiceServer) CreateSeveral(context.Context, *CreateSeveralFlangeType) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSeveral not implemented")
}
func (UnimplementedFlangeTypeServiceServer) Update(context.Context, *UpdateFlangeType) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedFlangeTypeServiceServer) Delete(context.Context, *DeleteFlangeType) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedFlangeTypeServiceServer) mustEmbedUnimplementedFlangeTypeServiceServer() {}

// UnsafeFlangeTypeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlangeTypeServiceServer will
// result in compilation errors.
type UnsafeFlangeTypeServiceServer interface {
	mustEmbedUnimplementedFlangeTypeServiceServer()
}

func RegisterFlangeTypeServiceServer(s grpc.ServiceRegistrar, srv FlangeTypeServiceServer) {
	s.RegisterService(&FlangeTypeService_ServiceDesc, srv)
}

func _FlangeTypeService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFlangeType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeTypeServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlangeTypeService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeTypeServiceServer).Get(ctx, req.(*GetFlangeType))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlangeTypeService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFlangeType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeTypeServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlangeTypeService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeTypeServiceServer).Create(ctx, req.(*CreateFlangeType))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlangeTypeService_CreateSeveral_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSeveralFlangeType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeTypeServiceServer).CreateSeveral(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlangeTypeService_CreateSeveral_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeTypeServiceServer).CreateSeveral(ctx, req.(*CreateSeveralFlangeType))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlangeTypeService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFlangeType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeTypeServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlangeTypeService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeTypeServiceServer).Update(ctx, req.(*UpdateFlangeType))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlangeTypeService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFlangeType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlangeTypeServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlangeTypeService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlangeTypeServiceServer).Delete(ctx, req.(*DeleteFlangeType))
	}
	return interceptor(ctx, in, info, handler)
}

// FlangeTypeService_ServiceDesc is the grpc.ServiceDesc for FlangeTypeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlangeTypeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flange_type_api.FlangeTypeService",
	HandlerType: (*FlangeTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _FlangeTypeService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _FlangeTypeService_Create_Handler,
		},
		{
			MethodName: "CreateSeveral",
			Handler:    _FlangeTypeService_CreateSeveral_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FlangeTypeService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FlangeTypeService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pro/flange_type_api.proto",
}
