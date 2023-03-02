// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: pro/mounting_api.proto

package mounting_api

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

// MountingServiceClient is the client API for MountingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MountingServiceClient interface {
	GetAll(ctx context.Context, in *GetAllMountings, opts ...grpc.CallOption) (*Mountings, error)
	Create(ctx context.Context, in *CreateMounting, opts ...grpc.CallOption) (*response_model.Response, error)
	CreateSeveral(ctx context.Context, in *CreateSeveralMounting, opts ...grpc.CallOption) (*response_model.Response, error)
	Update(ctx context.Context, in *UpdateMounting, opts ...grpc.CallOption) (*response_model.Response, error)
	Delete(ctx context.Context, in *DeleteMounting, opts ...grpc.CallOption) (*response_model.Response, error)
}

type mountingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMountingServiceClient(cc grpc.ClientConnInterface) MountingServiceClient {
	return &mountingServiceClient{cc}
}

func (c *mountingServiceClient) GetAll(ctx context.Context, in *GetAllMountings, opts ...grpc.CallOption) (*Mountings, error) {
	out := new(Mountings)
	err := c.cc.Invoke(ctx, "/mounting_api.MountingService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mountingServiceClient) Create(ctx context.Context, in *CreateMounting, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/mounting_api.MountingService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mountingServiceClient) CreateSeveral(ctx context.Context, in *CreateSeveralMounting, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/mounting_api.MountingService/CreateSeveral", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mountingServiceClient) Update(ctx context.Context, in *UpdateMounting, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/mounting_api.MountingService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mountingServiceClient) Delete(ctx context.Context, in *DeleteMounting, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/mounting_api.MountingService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MountingServiceServer is the server API for MountingService service.
// All implementations must embed UnimplementedMountingServiceServer
// for forward compatibility
type MountingServiceServer interface {
	GetAll(context.Context, *GetAllMountings) (*Mountings, error)
	Create(context.Context, *CreateMounting) (*response_model.Response, error)
	CreateSeveral(context.Context, *CreateSeveralMounting) (*response_model.Response, error)
	Update(context.Context, *UpdateMounting) (*response_model.Response, error)
	Delete(context.Context, *DeleteMounting) (*response_model.Response, error)
	mustEmbedUnimplementedMountingServiceServer()
}

// UnimplementedMountingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMountingServiceServer struct {
}

func (UnimplementedMountingServiceServer) GetAll(context.Context, *GetAllMountings) (*Mountings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedMountingServiceServer) Create(context.Context, *CreateMounting) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMountingServiceServer) CreateSeveral(context.Context, *CreateSeveralMounting) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSeveral not implemented")
}
func (UnimplementedMountingServiceServer) Update(context.Context, *UpdateMounting) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMountingServiceServer) Delete(context.Context, *DeleteMounting) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedMountingServiceServer) mustEmbedUnimplementedMountingServiceServer() {}

// UnsafeMountingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MountingServiceServer will
// result in compilation errors.
type UnsafeMountingServiceServer interface {
	mustEmbedUnimplementedMountingServiceServer()
}

func RegisterMountingServiceServer(s grpc.ServiceRegistrar, srv MountingServiceServer) {
	s.RegisterService(&MountingService_ServiceDesc, srv)
}

func _MountingService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllMountings)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MountingServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mounting_api.MountingService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MountingServiceServer).GetAll(ctx, req.(*GetAllMountings))
	}
	return interceptor(ctx, in, info, handler)
}

func _MountingService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMounting)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MountingServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mounting_api.MountingService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MountingServiceServer).Create(ctx, req.(*CreateMounting))
	}
	return interceptor(ctx, in, info, handler)
}

func _MountingService_CreateSeveral_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSeveralMounting)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MountingServiceServer).CreateSeveral(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mounting_api.MountingService/CreateSeveral",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MountingServiceServer).CreateSeveral(ctx, req.(*CreateSeveralMounting))
	}
	return interceptor(ctx, in, info, handler)
}

func _MountingService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMounting)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MountingServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mounting_api.MountingService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MountingServiceServer).Update(ctx, req.(*UpdateMounting))
	}
	return interceptor(ctx, in, info, handler)
}

func _MountingService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMounting)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MountingServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mounting_api.MountingService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MountingServiceServer).Delete(ctx, req.(*DeleteMounting))
	}
	return interceptor(ctx, in, info, handler)
}

// MountingService_ServiceDesc is the grpc.ServiceDesc for MountingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MountingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mounting_api.MountingService",
	HandlerType: (*MountingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _MountingService_GetAll_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _MountingService_Create_Handler,
		},
		{
			MethodName: "CreateSeveral",
			Handler:    _MountingService_CreateSeveral_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _MountingService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _MountingService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pro/mounting_api.proto",
}