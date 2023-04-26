// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: pro/putg_construction_type_api.proto

package putg_construction_api

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

// PutgConstructionServiceClient is the client API for PutgConstructionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PutgConstructionServiceClient interface {
	Get(ctx context.Context, in *GetPutgConstruction, opts ...grpc.CallOption) (*PutgConstruction, error)
	Create(ctx context.Context, in *CreatePutgConstruction, opts ...grpc.CallOption) (*response_model.Response, error)
	Update(ctx context.Context, in *UpdatePutgConstruction, opts ...grpc.CallOption) (*response_model.Response, error)
	Delete(ctx context.Context, in *DeletePutgConstruction, opts ...grpc.CallOption) (*response_model.Response, error)
}

type putgConstructionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPutgConstructionServiceClient(cc grpc.ClientConnInterface) PutgConstructionServiceClient {
	return &putgConstructionServiceClient{cc}
}

func (c *putgConstructionServiceClient) Get(ctx context.Context, in *GetPutgConstruction, opts ...grpc.CallOption) (*PutgConstruction, error) {
	out := new(PutgConstruction)
	err := c.cc.Invoke(ctx, "/putg_construction_api.PutgConstructionService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *putgConstructionServiceClient) Create(ctx context.Context, in *CreatePutgConstruction, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/putg_construction_api.PutgConstructionService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *putgConstructionServiceClient) Update(ctx context.Context, in *UpdatePutgConstruction, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/putg_construction_api.PutgConstructionService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *putgConstructionServiceClient) Delete(ctx context.Context, in *DeletePutgConstruction, opts ...grpc.CallOption) (*response_model.Response, error) {
	out := new(response_model.Response)
	err := c.cc.Invoke(ctx, "/putg_construction_api.PutgConstructionService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PutgConstructionServiceServer is the server API for PutgConstructionService service.
// All implementations must embed UnimplementedPutgConstructionServiceServer
// for forward compatibility
type PutgConstructionServiceServer interface {
	Get(context.Context, *GetPutgConstruction) (*PutgConstruction, error)
	Create(context.Context, *CreatePutgConstruction) (*response_model.Response, error)
	Update(context.Context, *UpdatePutgConstruction) (*response_model.Response, error)
	Delete(context.Context, *DeletePutgConstruction) (*response_model.Response, error)
	mustEmbedUnimplementedPutgConstructionServiceServer()
}

// UnimplementedPutgConstructionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPutgConstructionServiceServer struct {
}

func (UnimplementedPutgConstructionServiceServer) Get(context.Context, *GetPutgConstruction) (*PutgConstruction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPutgConstructionServiceServer) Create(context.Context, *CreatePutgConstruction) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPutgConstructionServiceServer) Update(context.Context, *UpdatePutgConstruction) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPutgConstructionServiceServer) Delete(context.Context, *DeletePutgConstruction) (*response_model.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPutgConstructionServiceServer) mustEmbedUnimplementedPutgConstructionServiceServer() {
}

// UnsafePutgConstructionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PutgConstructionServiceServer will
// result in compilation errors.
type UnsafePutgConstructionServiceServer interface {
	mustEmbedUnimplementedPutgConstructionServiceServer()
}

func RegisterPutgConstructionServiceServer(s grpc.ServiceRegistrar, srv PutgConstructionServiceServer) {
	s.RegisterService(&PutgConstructionService_ServiceDesc, srv)
}

func _PutgConstructionService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPutgConstruction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgConstructionServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_construction_api.PutgConstructionService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgConstructionServiceServer).Get(ctx, req.(*GetPutgConstruction))
	}
	return interceptor(ctx, in, info, handler)
}

func _PutgConstructionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePutgConstruction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgConstructionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_construction_api.PutgConstructionService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgConstructionServiceServer).Create(ctx, req.(*CreatePutgConstruction))
	}
	return interceptor(ctx, in, info, handler)
}

func _PutgConstructionService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePutgConstruction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgConstructionServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_construction_api.PutgConstructionService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgConstructionServiceServer).Update(ctx, req.(*UpdatePutgConstruction))
	}
	return interceptor(ctx, in, info, handler)
}

func _PutgConstructionService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePutgConstruction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgConstructionServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_construction_api.PutgConstructionService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgConstructionServiceServer).Delete(ctx, req.(*DeletePutgConstruction))
	}
	return interceptor(ctx, in, info, handler)
}

// PutgConstructionService_ServiceDesc is the grpc.ServiceDesc for PutgConstructionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PutgConstructionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "putg_construction_api.PutgConstructionService",
	HandlerType: (*PutgConstructionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PutgConstructionService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _PutgConstructionService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PutgConstructionService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PutgConstructionService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pro/putg_construction_type_api.proto",
}
