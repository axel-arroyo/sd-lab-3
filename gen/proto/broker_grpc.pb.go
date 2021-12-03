// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// BrokerClient is the client API for Broker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrokerClient interface {
	AddCity(ctx context.Context, in *AddCityRequest, opts ...grpc.CallOption) (*AddCityBrokerResponse, error)
	DeleteCity(ctx context.Context, in *DeleteCityRequest, opts ...grpc.CallOption) (*DeleteCityBrokerResponse, error)
	UpdateName(ctx context.Context, in *UpdateNameRequest, opts ...grpc.CallOption) (*UpdateNameBrokerResponse, error)
	UpdateNumber(ctx context.Context, in *UpdateNumberRequest, opts ...grpc.CallOption) (*UpdateNumberBrokerResponse, error)
	GetNumberRebeldes(ctx context.Context, in *GetNumberRebeldesRequest, opts ...grpc.CallOption) (*GetNumberRebeldesResponse, error)
}

type brokerClient struct {
	cc grpc.ClientConnInterface
}

func NewBrokerClient(cc grpc.ClientConnInterface) BrokerClient {
	return &brokerClient{cc}
}

func (c *brokerClient) AddCity(ctx context.Context, in *AddCityRequest, opts ...grpc.CallOption) (*AddCityBrokerResponse, error) {
	out := new(AddCityBrokerResponse)
	err := c.cc.Invoke(ctx, "/grpc.Broker/AddCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerClient) DeleteCity(ctx context.Context, in *DeleteCityRequest, opts ...grpc.CallOption) (*DeleteCityBrokerResponse, error) {
	out := new(DeleteCityBrokerResponse)
	err := c.cc.Invoke(ctx, "/grpc.Broker/DeleteCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerClient) UpdateName(ctx context.Context, in *UpdateNameRequest, opts ...grpc.CallOption) (*UpdateNameBrokerResponse, error) {
	out := new(UpdateNameBrokerResponse)
	err := c.cc.Invoke(ctx, "/grpc.Broker/UpdateName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerClient) UpdateNumber(ctx context.Context, in *UpdateNumberRequest, opts ...grpc.CallOption) (*UpdateNumberBrokerResponse, error) {
	out := new(UpdateNumberBrokerResponse)
	err := c.cc.Invoke(ctx, "/grpc.Broker/UpdateNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerClient) GetNumberRebeldes(ctx context.Context, in *GetNumberRebeldesRequest, opts ...grpc.CallOption) (*GetNumberRebeldesResponse, error) {
	out := new(GetNumberRebeldesResponse)
	err := c.cc.Invoke(ctx, "/grpc.Broker/GetNumberRebeldes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrokerServer is the server API for Broker service.
// All implementations must embed UnimplementedBrokerServer
// for forward compatibility
type BrokerServer interface {
	AddCity(context.Context, *AddCityRequest) (*AddCityBrokerResponse, error)
	DeleteCity(context.Context, *DeleteCityRequest) (*DeleteCityBrokerResponse, error)
	UpdateName(context.Context, *UpdateNameRequest) (*UpdateNameBrokerResponse, error)
	UpdateNumber(context.Context, *UpdateNumberRequest) (*UpdateNumberBrokerResponse, error)
	GetNumberRebeldes(context.Context, *GetNumberRebeldesRequest) (*GetNumberRebeldesResponse, error)
	mustEmbedUnimplementedBrokerServer()
}

// UnimplementedBrokerServer must be embedded to have forward compatible implementations.
type UnimplementedBrokerServer struct {
}

func (UnimplementedBrokerServer) AddCity(context.Context, *AddCityRequest) (*AddCityBrokerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCity not implemented")
}
func (UnimplementedBrokerServer) DeleteCity(context.Context, *DeleteCityRequest) (*DeleteCityBrokerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCity not implemented")
}
func (UnimplementedBrokerServer) UpdateName(context.Context, *UpdateNameRequest) (*UpdateNameBrokerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateName not implemented")
}
func (UnimplementedBrokerServer) UpdateNumber(context.Context, *UpdateNumberRequest) (*UpdateNumberBrokerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNumber not implemented")
}
func (UnimplementedBrokerServer) GetNumberRebeldes(context.Context, *GetNumberRebeldesRequest) (*GetNumberRebeldesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNumberRebeldes not implemented")
}
func (UnimplementedBrokerServer) mustEmbedUnimplementedBrokerServer() {}

// UnsafeBrokerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrokerServer will
// result in compilation errors.
type UnsafeBrokerServer interface {
	mustEmbedUnimplementedBrokerServer()
}

func RegisterBrokerServer(s grpc.ServiceRegistrar, srv BrokerServer) {
	s.RegisterService(&Broker_ServiceDesc, srv)
}

func _Broker_AddCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).AddCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Broker/AddCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).AddCity(ctx, req.(*AddCityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broker_DeleteCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).DeleteCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Broker/DeleteCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).DeleteCity(ctx, req.(*DeleteCityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broker_UpdateName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).UpdateName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Broker/UpdateName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).UpdateName(ctx, req.(*UpdateNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broker_UpdateNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).UpdateNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Broker/UpdateNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).UpdateNumber(ctx, req.(*UpdateNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broker_GetNumberRebeldes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNumberRebeldesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).GetNumberRebeldes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Broker/GetNumberRebeldes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).GetNumberRebeldes(ctx, req.(*GetNumberRebeldesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Broker_ServiceDesc is the grpc.ServiceDesc for Broker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Broker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Broker",
	HandlerType: (*BrokerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCity",
			Handler:    _Broker_AddCity_Handler,
		},
		{
			MethodName: "DeleteCity",
			Handler:    _Broker_DeleteCity_Handler,
		},
		{
			MethodName: "UpdateName",
			Handler:    _Broker_UpdateName_Handler,
		},
		{
			MethodName: "UpdateNumber",
			Handler:    _Broker_UpdateNumber_Handler,
		},
		{
			MethodName: "GetNumberRebeldes",
			Handler:    _Broker_GetNumberRebeldes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "broker.proto",
}
