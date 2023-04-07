// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: pro/snp_size_api.proto

package snp_size_api

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

// SnpSizeServiceClient is the client API for SnpSizeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SnpSizeServiceClient interface {
	Get(ctx context.Context, in *GetSnpSize, opts ...grpc.CallOption) (*SnpSize, error)
	Create(ctx context.Context, in *CreateSnpSize, opts ...grpc.CallOption) (*response_model.Response, error)
	CreateSeveral(ctx context.Context, in *CreateSeveralSnpSize, opts ...grpc.CallOption) (*response_model.Response, error)
	Update(ctx context.Context, in *UpdateSnpSize, opts ...grpc.CallOption) (*response_model.Response, error)
	Delete(ctx context.Context, in *DeleteSnpSize, opts ...grpc.CallOption) (*response_model.Response, error)
}

type snpSizeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSnpSizeServiceClient(cc grpc.ClientConnInterface) SnpSizeServiceClient {
	return &snpSizeServiceClient{cc}
}

func (c *snpSizeServiceClient) Get(ctx context.Context, in *GetSnpSize, opts ...grpc.CallOption) (*SnpSize, error) {
	out := new(SnpSize)
	err := c.cc.Invoke(ctx, "/snp_size_api.SnpSizeService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snpSizeServiceClient) Create(ctx context.Context, in *CreateSnpSize, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/snp_size_api.SnpSizeService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snpSizeServiceClient) CreateSeveral(ctx context.Context, in *CreateSeveralSnpSize, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/snp_size_api.SnpSizeService/CreateSeveral", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snpSizeServiceClient) Update(ctx context.Context, in *UpdateSnpSize, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/snp_size_api.SnpSizeService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snpSizeServiceClient) Delete(ctx context.Context, in *DeleteSnpSize, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/snp_size_api.SnpSizeService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SnpSizeServiceServer is the server API for SnpSizeService service.
// All implementations must embed UnimplementedSnpSizeServiceServer
// for forward compatibility
type SnpSizeServiceServer interface {
	Get(context.Context, *GetSnpSize) (*SnpSize, error)
	Create(context.Context, *CreateSnpSize) (*response_model.Response, error)
	CreateSeveral(context.Context, *CreateSeveralSnpSize) (*response_model.Response, error)
	Update(context.Context, *UpdateSnpSize) (*response_model.Response, error)
	Delete(context.Context, *DeleteSnpSize) (*response_model.Response, error)
	mustEmbedUnimplementedSnpSizeServiceServer()
}

// UnimplementedSnpSizeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSnpSizeServiceServer struct {
}

func (UnimplementedSnpSizeServiceServer) Get(context.Context, *GetSnpSize) (*SnpSize, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSnpSizeServiceServer) Create(context.Context, *CreateSnpSize) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSnpSizeServiceServer) CreateSeveral(context.Context, *CreateSeveralSnpSize) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSeveral not implemented")
}
func (UnimplementedSnpSizeServiceServer) Update(context.Context, *UpdateSnpSize) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSnpSizeServiceServer) Delete(context.Context, *DeleteSnpSize) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSnpSizeServiceServer) mustEmbedUnimplementedSnpSizeServiceServer() {}

// UnsafeSnpSizeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SnpSizeServiceServer will
// result in compilation errors.
type UnsafeSnpSizeServiceServer interface {
	mustEmbedUnimplementedSnpSizeServiceServer()
}

func RegisterSnpSizeServiceServer(s grpc.ServiceRegistrar, srv SnpSizeServiceServer) {
	s.RegisterService(&SnpSizeService_ServiceDesc, srv)
}

func _SnpSizeService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSnpSize)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnpSizeServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snp_size_api.SnpSizeService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnpSizeServiceServer).Get(ctx, req.(*GetSnpSize))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnpSizeService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSnpSize)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnpSizeServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snp_size_api.SnpSizeService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnpSizeServiceServer).Create(ctx, req.(*CreateSnpSize))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnpSizeService_CreateSeveral_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSeveralSnpSize)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnpSizeServiceServer).CreateSeveral(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snp_size_api.SnpSizeService/CreateSeveral",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnpSizeServiceServer).CreateSeveral(ctx, req.(*CreateSeveralSnpSize))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnpSizeService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSnpSize)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnpSizeServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snp_size_api.SnpSizeService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnpSizeServiceServer).Update(ctx, req.(*UpdateSnpSize))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnpSizeService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSnpSize)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnpSizeServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snp_size_api.SnpSizeService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnpSizeServiceServer).Delete(ctx, req.(*DeleteSnpSize))
	}
	return interceptor(ctx, in, info, handler)
}

// SnpSizeService_ServiceDesc is the grpc.ServiceDesc for SnpSizeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SnpSizeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "snp_size_api.SnpSizeService",
	HandlerType: (*SnpSizeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SnpSizeService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _SnpSizeService_Create_Handler,
		},
		{
			MethodName: "CreateSeveral",
			Handler:    _SnpSizeService_CreateSeveral_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SnpSizeService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SnpSizeService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pro/snp_size_api.proto",
}
