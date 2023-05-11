// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: pro/putg_material_api.proto

package putg_material_api

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

// PutgMaterialServiceClient is the client API for PutgMaterialService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PutgMaterialServiceClient interface {
	Get(ctx context.Context, in *GetPutgMaterial, opts ...grpc.CallOption) (*PutgMaterials, error)
	Create(ctx context.Context, in *CreatePutgMaterial, opts ...grpc.CallOption) (*response_model.Response, error)
	Update(ctx context.Context, in *UpdatePutgMaterial, opts ...grpc.CallOption) (*response_model.Response, error)
	Delete(ctx context.Context, in *DeletePutgMaterial, opts ...grpc.CallOption) (*response_model.Response, error)
}

type putgMaterialServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPutgMaterialServiceClient(cc grpc.ClientConnInterface) PutgMaterialServiceClient {
	return &putgMaterialServiceClient{cc}
}

func (c *putgMaterialServiceClient) Get(ctx context.Context, in *GetPutgMaterial, opts ...grpc.CallOption) (*PutgMaterials, error) {
	out := new(PutgMaterials)
	err := c.cc.Invoke(ctx, "/putg_material_api.PutgMaterialService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *putgMaterialServiceClient) Create(ctx context.Context, in *CreatePutgMaterial, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/putg_material_api.PutgMaterialService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *putgMaterialServiceClient) Update(ctx context.Context, in *UpdatePutgMaterial, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/putg_material_api.PutgMaterialService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *putgMaterialServiceClient) Delete(ctx context.Context, in *DeletePutgMaterial, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/putg_material_api.PutgMaterialService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PutgMaterialServiceServer is the server API for PutgMaterialService service.
// All implementations must embed UnimplementedPutgMaterialServiceServer
// for forward compatibility
type PutgMaterialServiceServer interface {
	Get(context.Context, *GetPutgMaterial) (*PutgMaterials, error)
	Create(context.Context, *CreatePutgMaterial) (*response_model.Response, error)
	Update(context.Context, *UpdatePutgMaterial) (*response_model.Response, error)
	Delete(context.Context, *DeletePutgMaterial) (*response_model.Response, error)
	mustEmbedUnimplementedPutgMaterialServiceServer()
}

// UnimplementedPutgMaterialServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPutgMaterialServiceServer struct {
}

func (UnimplementedPutgMaterialServiceServer) Get(context.Context, *GetPutgMaterial) (*PutgMaterials, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPutgMaterialServiceServer) Create(context.Context, *CreatePutgMaterial) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPutgMaterialServiceServer) Update(context.Context, *UpdatePutgMaterial) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPutgMaterialServiceServer) Delete(context.Context, *DeletePutgMaterial) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPutgMaterialServiceServer) mustEmbedUnimplementedPutgMaterialServiceServer() {}

// UnsafePutgMaterialServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PutgMaterialServiceServer will
// result in compilation errors.
type UnsafePutgMaterialServiceServer interface {
	mustEmbedUnimplementedPutgMaterialServiceServer()
}

func RegisterPutgMaterialServiceServer(s grpc.ServiceRegistrar, srv PutgMaterialServiceServer) {
	s.RegisterService(&PutgMaterialService_ServiceDesc, srv)
}

func _PutgMaterialService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPutgMaterial)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgMaterialServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_material_api.PutgMaterialService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgMaterialServiceServer).Get(ctx, req.(*GetPutgMaterial))
	}
	return interceptor(ctx, in, info, handler)
}

func _PutgMaterialService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePutgMaterial)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgMaterialServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_material_api.PutgMaterialService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgMaterialServiceServer).Create(ctx, req.(*CreatePutgMaterial))
	}
	return interceptor(ctx, in, info, handler)
}

func _PutgMaterialService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePutgMaterial)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgMaterialServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_material_api.PutgMaterialService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgMaterialServiceServer).Update(ctx, req.(*UpdatePutgMaterial))
	}
	return interceptor(ctx, in, info, handler)
}

func _PutgMaterialService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePutgMaterial)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgMaterialServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_material_api.PutgMaterialService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgMaterialServiceServer).Delete(ctx, req.(*DeletePutgMaterial))
	}
	return interceptor(ctx, in, info, handler)
}

// PutgMaterialService_ServiceDesc is the grpc.ServiceDesc for PutgMaterialService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PutgMaterialService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "putg_material_api.PutgMaterialService",
	HandlerType: (*PutgMaterialServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PutgMaterialService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _PutgMaterialService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PutgMaterialService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PutgMaterialService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pro/putg_material_api.proto",
}