// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: moment/calc_api.proto

package calc_api

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

// CalcServiceClient is the client API for CalcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalcServiceClient interface {
	CalculateFlange(ctx context.Context, in *FlangeRequest, opts ...grpc.CallOption) (*FlangeResponse, error)
	CalculateCap(ctx context.Context, in *CapRequest, opts ...grpc.CallOption) (*CapResponse, error)
	CalculateFloat(ctx context.Context, in *FloatRequest, opts ...grpc.CallOption) (*FloatResponse, error)
	CalculateDevCooling(ctx context.Context, in *DevCoolingRequest, opts ...grpc.CallOption) (*DevCoolingResponse, error)
}

type calcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalcServiceClient(cc grpc.ClientConnInterface) CalcServiceClient {
	return &calcServiceClient{cc}
}

func (c *calcServiceClient) CalculateFlange(ctx context.Context, in *FlangeRequest, opts ...grpc.CallOption) (*FlangeResponse, error) {
	out := new(FlangeResponse)
	err := c.cc.Invoke(ctx, "/calc_api.CalcService/CalculateFlange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) CalculateCap(ctx context.Context, in *CapRequest, opts ...grpc.CallOption) (*CapResponse, error) {
	out := new(CapResponse)
	err := c.cc.Invoke(ctx, "/calc_api.CalcService/CalculateCap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) CalculateFloat(ctx context.Context, in *FloatRequest, opts ...grpc.CallOption) (*FloatResponse, error) {
	out := new(FloatResponse)
	err := c.cc.Invoke(ctx, "/calc_api.CalcService/CalculateFloat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) CalculateDevCooling(ctx context.Context, in *DevCoolingRequest, opts ...grpc.CallOption) (*DevCoolingResponse, error) {
	out := new(DevCoolingResponse)
	err := c.cc.Invoke(ctx, "/calc_api.CalcService/CalculateDevCooling", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalcServiceServer is the server API for CalcService service.
// All implementations must embed UnimplementedCalcServiceServer
// for forward compatibility
type CalcServiceServer interface {
	CalculateFlange(context.Context, *FlangeRequest) (*FlangeResponse, error)
	CalculateCap(context.Context, *CapRequest) (*CapResponse, error)
	CalculateFloat(context.Context, *FloatRequest) (*FloatResponse, error)
	CalculateDevCooling(context.Context, *DevCoolingRequest) (*DevCoolingResponse, error)
	mustEmbedUnimplementedCalcServiceServer()
}

// UnimplementedCalcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalcServiceServer struct {
}

func (UnimplementedCalcServiceServer) CalculateFlange(context.Context, *FlangeRequest) (*FlangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateFlange not implemented")
}
func (UnimplementedCalcServiceServer) CalculateCap(context.Context, *CapRequest) (*CapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateCap not implemented")
}
func (UnimplementedCalcServiceServer) CalculateFloat(context.Context, *FloatRequest) (*FloatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateFloat not implemented")
}
func (UnimplementedCalcServiceServer) CalculateDevCooling(context.Context, *DevCoolingRequest) (*DevCoolingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateDevCooling not implemented")
}
func (UnimplementedCalcServiceServer) mustEmbedUnimplementedCalcServiceServer() {}

// UnsafeCalcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalcServiceServer will
// result in compilation errors.
type UnsafeCalcServiceServer interface {
	mustEmbedUnimplementedCalcServiceServer()
}

func RegisterCalcServiceServer(s grpc.ServiceRegistrar, srv CalcServiceServer) {
	s.RegisterService(&CalcService_ServiceDesc, srv)
}

func _CalcService_CalculateFlange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).CalculateFlange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc_api.CalcService/CalculateFlange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).CalculateFlange(ctx, req.(*FlangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_CalculateCap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).CalculateCap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc_api.CalcService/CalculateCap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).CalculateCap(ctx, req.(*CapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_CalculateFloat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FloatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).CalculateFloat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc_api.CalcService/CalculateFloat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).CalculateFloat(ctx, req.(*FloatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_CalculateDevCooling_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DevCoolingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).CalculateDevCooling(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc_api.CalcService/CalculateDevCooling",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).CalculateDevCooling(ctx, req.(*DevCoolingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalcService_ServiceDesc is the grpc.ServiceDesc for CalcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calc_api.CalcService",
	HandlerType: (*CalcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalculateFlange",
			Handler:    _CalcService_CalculateFlange_Handler,
		},
		{
			MethodName: "CalculateCap",
			Handler:    _CalcService_CalculateCap_Handler,
		},
		{
			MethodName: "CalculateFloat",
			Handler:    _CalcService_CalculateFloat_Handler,
		},
		{
			MethodName: "CalculateDevCooling",
			Handler:    _CalcService_CalculateDevCooling_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moment/calc_api.proto",
}
