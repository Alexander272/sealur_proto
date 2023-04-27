// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: pro/putg_api.proto

package putg_api

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

// PutgDataServiceClient is the client API for PutgDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PutgDataServiceClient interface {
	Get(ctx context.Context, in *GetPutg, opts ...grpc.CallOption) (*Putg, error)
	GetData(ctx context.Context, in *GetPutgData, opts ...grpc.CallOption) (*PutgData, error)
	GetBase(ctx context.Context, in *GetPutgBase, opts ...grpc.CallOption) (*PutgBase, error)
}

type putgDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPutgDataServiceClient(cc grpc.ClientConnInterface) PutgDataServiceClient {
	return &putgDataServiceClient{cc}
}

func (c *putgDataServiceClient) Get(ctx context.Context, in *GetPutg, opts ...grpc.CallOption) (*Putg, error) {
	out := new(Putg)
	err := c.cc.Invoke(ctx, "/putg_api.PutgDataService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *putgDataServiceClient) GetData(ctx context.Context, in *GetPutgData, opts ...grpc.CallOption) (*PutgData, error) {
	out := new(PutgData)
	err := c.cc.Invoke(ctx, "/putg_api.PutgDataService/GetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *putgDataServiceClient) GetBase(ctx context.Context, in *GetPutgBase, opts ...grpc.CallOption) (*PutgBase, error) {
	out := new(PutgBase)
	err := c.cc.Invoke(ctx, "/putg_api.PutgDataService/GetBase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PutgDataServiceServer is the server API for PutgDataService service.
// All implementations must embed UnimplementedPutgDataServiceServer
// for forward compatibility
type PutgDataServiceServer interface {
	Get(context.Context, *GetPutg) (*Putg, error)
	GetData(context.Context, *GetPutgData) (*PutgData, error)
	GetBase(context.Context, *GetPutgBase) (*PutgBase, error)
	mustEmbedUnimplementedPutgDataServiceServer()
}

// UnimplementedPutgDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPutgDataServiceServer struct {
}

func (UnimplementedPutgDataServiceServer) Get(context.Context, *GetPutg) (*Putg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPutgDataServiceServer) GetData(context.Context, *GetPutgData) (*PutgData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (UnimplementedPutgDataServiceServer) GetBase(context.Context, *GetPutgBase) (*PutgBase, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBase not implemented")
}
func (UnimplementedPutgDataServiceServer) mustEmbedUnimplementedPutgDataServiceServer() {}

// UnsafePutgDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PutgDataServiceServer will
// result in compilation errors.
type UnsafePutgDataServiceServer interface {
	mustEmbedUnimplementedPutgDataServiceServer()
}

func RegisterPutgDataServiceServer(s grpc.ServiceRegistrar, srv PutgDataServiceServer) {
	s.RegisterService(&PutgDataService_ServiceDesc, srv)
}

func _PutgDataService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPutg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgDataServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_api.PutgDataService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgDataServiceServer).Get(ctx, req.(*GetPutg))
	}
	return interceptor(ctx, in, info, handler)
}

func _PutgDataService_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPutgData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgDataServiceServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_api.PutgDataService/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgDataServiceServer).GetData(ctx, req.(*GetPutgData))
	}
	return interceptor(ctx, in, info, handler)
}

func _PutgDataService_GetBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPutgBase)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PutgDataServiceServer).GetBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/putg_api.PutgDataService/GetBase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PutgDataServiceServer).GetBase(ctx, req.(*GetPutgBase))
	}
	return interceptor(ctx, in, info, handler)
}

// PutgDataService_ServiceDesc is the grpc.ServiceDesc for PutgDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PutgDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "putg_api.PutgDataService",
	HandlerType: (*PutgDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PutgDataService_Get_Handler,
		},
		{
			MethodName: "GetData",
			Handler:    _PutgDataService_GetData_Handler,
		},
		{
			MethodName: "GetBase",
			Handler:    _PutgDataService_GetBase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pro/putg_api.proto",
}
