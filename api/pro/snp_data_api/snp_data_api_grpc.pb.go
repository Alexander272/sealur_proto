// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: pro/snp_data_api.proto

package snp_data_api

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
	SnpDataService_Get_FullMethodName    = "/snp_data_api.SnpDataService/Get"
	SnpDataService_Create_FullMethodName = "/snp_data_api.SnpDataService/Create"
	SnpDataService_Update_FullMethodName = "/snp_data_api.SnpDataService/Update"
	SnpDataService_Delete_FullMethodName = "/snp_data_api.SnpDataService/Delete"
)

// SnpDataServiceClient is the client API for SnpDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SnpDataServiceClient interface {
	Get(ctx context.Context, in *GetSnpData, opts ...grpc.CallOption) (*SnpData, error)
	Create(ctx context.Context, in *CreateSnpData, opts ...grpc.CallOption) (*response_model.Response, error)
	Update(ctx context.Context, in *UpdateSnpData, opts ...grpc.CallOption) (*response_model.Response, error)
	Delete(ctx context.Context, in *DeleteSnpData, opts ...grpc.CallOption) (*response_model.Response, error)
}

type snpDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSnpDataServiceClient(cc grpc.ClientConnInterface) SnpDataServiceClient {
	return &snpDataServiceClient{cc}
}

func (c *snpDataServiceClient) Get(ctx context.Context, in *GetSnpData, opts ...grpc.CallOption) (*SnpData, error) {
	out := new(SnpData)
	err := c.cc.Invoke(ctx, SnpDataService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snpDataServiceClient) Create(ctx context.Context, in *CreateSnpData, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, SnpDataService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snpDataServiceClient) Update(ctx context.Context, in *UpdateSnpData, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, SnpDataService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snpDataServiceClient) Delete(ctx context.Context, in *DeleteSnpData, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, SnpDataService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SnpDataServiceServer is the server API for SnpDataService service.
// All implementations must embed UnimplementedSnpDataServiceServer
// for forward compatibility
type SnpDataServiceServer interface {
	Get(context.Context, *GetSnpData) (*SnpData, error)
	Create(context.Context, *CreateSnpData) (*response_model.Response, error)
	Update(context.Context, *UpdateSnpData) (*response_model.Response, error)
	Delete(context.Context, *DeleteSnpData) (*response_model.Response, error)
	mustEmbedUnimplementedSnpDataServiceServer()
}

// UnimplementedSnpDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSnpDataServiceServer struct {
}

func (UnimplementedSnpDataServiceServer) Get(context.Context, *GetSnpData) (*SnpData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSnpDataServiceServer) Create(context.Context, *CreateSnpData) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSnpDataServiceServer) Update(context.Context, *UpdateSnpData) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSnpDataServiceServer) Delete(context.Context, *DeleteSnpData) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSnpDataServiceServer) mustEmbedUnimplementedSnpDataServiceServer() {}

// UnsafeSnpDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SnpDataServiceServer will
// result in compilation errors.
type UnsafeSnpDataServiceServer interface {
	mustEmbedUnimplementedSnpDataServiceServer()
}

func RegisterSnpDataServiceServer(s grpc.ServiceRegistrar, srv SnpDataServiceServer) {
	s.RegisterService(&SnpDataService_ServiceDesc, srv)
}

func _SnpDataService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSnpData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnpDataServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnpDataService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnpDataServiceServer).Get(ctx, req.(*GetSnpData))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnpDataService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSnpData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnpDataServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnpDataService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnpDataServiceServer).Create(ctx, req.(*CreateSnpData))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnpDataService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSnpData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnpDataServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnpDataService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnpDataServiceServer).Update(ctx, req.(*UpdateSnpData))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnpDataService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSnpData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnpDataServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnpDataService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnpDataServiceServer).Delete(ctx, req.(*DeleteSnpData))
	}
	return interceptor(ctx, in, info, handler)
}

// SnpDataService_ServiceDesc is the grpc.ServiceDesc for SnpDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SnpDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "snp_data_api.SnpDataService",
	HandlerType: (*SnpDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SnpDataService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _SnpDataService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SnpDataService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SnpDataService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pro/snp_data_api.proto",
}
