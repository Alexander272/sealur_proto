// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: pro/putg_construction_base_api.proto

package putg_base_construction_api

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
	PutgBaseConstructionService_Get_FullMethodName    = "/putg_base_construction_api.PutgBaseConstructionService/Get"
	PutgBaseConstructionService_Create_FullMethodName = "/putg_base_construction_api.PutgBaseConstructionService/Create"
	PutgBaseConstructionService_Update_FullMethodName = "/putg_base_construction_api.PutgBaseConstructionService/Update"
	PutgBaseConstructionService_Delete_FullMethodName = "/putg_base_construction_api.PutgBaseConstructionService/Delete"
)

// PutgBaseConstructionServiceClient is the client API for PutgBaseConstructionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PutgBaseConstructionServiceClient interface {
	Get(ctx context.Context, in *GetPutgBaseConstruction, opts ...grpc.CallOption) (*PutgBaseConstruction, error)
	Create(ctx context.Context, in *CreatePutgBaseConstruction, opts ...grpc.CallOption) (*response_model.Response, error)
	Update(ctx context.Context, in *UpdatePutgBaseConstruction, opts ...grpc.CallOption) (*response_model.Response, error)
	Delete(ctx context.Context, in *DeletePutgBaseConstruction, opts ...grpc.CallOption) (*response_model.Response, error)
}

type putgBaseConstructionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPutgBaseConstructionServiceClient(cc grpc.ClientConnInterface) PutgBaseConstructionServiceClient {
	return &putgBaseConstructionServiceClient{cc}
}

func (c *putgBaseConstructionServiceClient) Get(ctx context.Context, in *GetPutgBaseConstruction, opts ...grpc.CallOption) (*PutgBaseConstruction, error) {
	out := new(PutgBaseConstruction)
	err := c.cc.Invoke(ctx, PutgBaseConstructionService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *putgBaseConstructionServiceClient) Create(ctx context.Context, in *CreatePutgBaseConstruction, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, PutgBaseConstructionService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *putgBaseConstructionServiceClient) Update(ctx context.Context, in *UpdatePutgBaseConstruction, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, PutgBaseConstructionService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *putgBaseConstructionServiceClient) Delete(ctx context.Context, in *DeletePutgBaseConstruction, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, PutgBaseConstructionService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PutgBaseConstructionServiceServer is the server API for PutgBaseConstructionService service.
// All implementations must embed UnimplementedPutgBaseConstructionServiceServer
// for forward compatibility
type PutgBaseConstructionServiceServer interface {
	Get(context.Context, *GetPutgBaseConstruction) (*PutgBaseConstruction, error)
	Create(context.Context, *CreatePutgBaseConstruction) (*response_model.Response, error)
	Update(context.Context, *UpdatePutgBaseConstruction) (*response_model.Response, error)
	Delete(context.Context, *DeletePutgBaseConstruction) (*response_model.Response, error)
	mustEmbedUnimplementedPutgBaseConstructionServiceServer()
}

// UnimplementedPutgBaseConstructionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPutgBaseConstructionServiceServer struct {
}

func (UnimplementedPutgBaseConstructionServiceServer) Get(context.Context, *GetPutgBaseConstruction) (*PutgBaseConstruction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPutgBaseConstructionServiceServer) Create(context.Context, *CreatePutgBaseConstruction) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPutgBaseConstructionServiceServer) Update(context.Context, *UpdatePutgBaseConstruction) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPutgBaseConstructionServiceServer) Delete(context.Context, *DeletePutgBaseConstruction) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPutgBaseConstructionServiceServer) mustEmbedUnimplementedPutgBaseConstructionServiceServer() {
}

// UnsafePutgBaseConstructionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PutgBaseConstructionServiceServer will
// result in compilation errors.
type UnsafePutgBaseConstructionServiceServer interface {
	mustEmbedUnimplementedPutgBaseConstructionServiceServer()
}

func RegisterPutgBaseConstructionServiceServer(s grpc.ServiceRegistrar, srv PutgBaseConstructionServiceServer) {
	s.RegisterService(&PutgBaseConstructionService_ServiceDesc, srv)
}

func _PutgBaseConstructionService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPutgBaseConstruction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgBaseConstructionServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PutgBaseConstructionService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgBaseConstructionServiceServer).Get(ctx, req.(*GetPutgBaseConstruction))
	}
	return interceptor(ctx, in, info, handler)
}

func _PutgBaseConstructionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePutgBaseConstruction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgBaseConstructionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PutgBaseConstructionService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgBaseConstructionServiceServer).Create(ctx, req.(*CreatePutgBaseConstruction))
	}
	return interceptor(ctx, in, info, handler)
}

func _PutgBaseConstructionService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePutgBaseConstruction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgBaseConstructionServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PutgBaseConstructionService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgBaseConstructionServiceServer).Update(ctx, req.(*UpdatePutgBaseConstruction))
	}
	return interceptor(ctx, in, info, handler)
}

func _PutgBaseConstructionService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePutgBaseConstruction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgBaseConstructionServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PutgBaseConstructionService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgBaseConstructionServiceServer).Delete(ctx, req.(*DeletePutgBaseConstruction))
	}
	return interceptor(ctx, in, info, handler)
}

// PutgBaseConstructionService_ServiceDesc is the grpc.ServiceDesc for PutgBaseConstructionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PutgBaseConstructionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "putg_base_construction_api.PutgBaseConstructionService",
	HandlerType: (*PutgBaseConstructionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PutgBaseConstructionService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _PutgBaseConstructionService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PutgBaseConstructionService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PutgBaseConstructionService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pro/putg_construction_base_api.proto",
}
