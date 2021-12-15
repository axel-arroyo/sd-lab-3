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

// FulcrumClient is the client API for Fulcrum service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FulcrumClient interface {
	IsAvailable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// functions
	AddCity(ctx context.Context, in *AddCityRequest, opts ...grpc.CallOption) (*VectorClock, error)
	DeleteCity(ctx context.Context, in *DeleteCityRequest, opts ...grpc.CallOption) (*VectorClock, error)
	UpdateName(ctx context.Context, in *UpdateNameRequest, opts ...grpc.CallOption) (*VectorClock, error)
	UpdateNumber(ctx context.Context, in *UpdateNumberRequest, opts ...grpc.CallOption) (*VectorClock, error)
	GetNumberRebeldesFulcrum(ctx context.Context, in *GetNumberRebeldesRequest, opts ...grpc.CallOption) (*GetNumberRebeldesResponse, error)
	// Merge
	VectorClockMerge(ctx context.Context, in *VectorClock, opts ...grpc.CallOption) (*Empty, error)
	Merge(ctx context.Context, opts ...grpc.CallOption) (Fulcrum_MergeClient, error)
	MergeFulcrums(ctx context.Context, opts ...grpc.CallOption) (Fulcrum_MergeFulcrumsClient, error)
}

type fulcrumClient struct {
	cc grpc.ClientConnInterface
}

func NewFulcrumClient(cc grpc.ClientConnInterface) FulcrumClient {
	return &fulcrumClient{cc}
}

func (c *fulcrumClient) IsAvailable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/grpc.Fulcrum/IsAvailable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumClient) AddCity(ctx context.Context, in *AddCityRequest, opts ...grpc.CallOption) (*VectorClock, error) {
	out := new(VectorClock)
	err := c.cc.Invoke(ctx, "/grpc.Fulcrum/AddCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumClient) DeleteCity(ctx context.Context, in *DeleteCityRequest, opts ...grpc.CallOption) (*VectorClock, error) {
	out := new(VectorClock)
	err := c.cc.Invoke(ctx, "/grpc.Fulcrum/DeleteCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumClient) UpdateName(ctx context.Context, in *UpdateNameRequest, opts ...grpc.CallOption) (*VectorClock, error) {
	out := new(VectorClock)
	err := c.cc.Invoke(ctx, "/grpc.Fulcrum/UpdateName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumClient) UpdateNumber(ctx context.Context, in *UpdateNumberRequest, opts ...grpc.CallOption) (*VectorClock, error) {
	out := new(VectorClock)
	err := c.cc.Invoke(ctx, "/grpc.Fulcrum/UpdateNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumClient) GetNumberRebeldesFulcrum(ctx context.Context, in *GetNumberRebeldesRequest, opts ...grpc.CallOption) (*GetNumberRebeldesResponse, error) {
	out := new(GetNumberRebeldesResponse)
	err := c.cc.Invoke(ctx, "/grpc.Fulcrum/GetNumberRebeldesFulcrum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumClient) VectorClockMerge(ctx context.Context, in *VectorClock, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/grpc.Fulcrum/VectorClockMerge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumClient) Merge(ctx context.Context, opts ...grpc.CallOption) (Fulcrum_MergeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Fulcrum_ServiceDesc.Streams[0], "/grpc.Fulcrum/Merge", opts...)
	if err != nil {
		return nil, err
	}
	x := &fulcrumMergeClient{stream}
	return x, nil
}

type Fulcrum_MergeClient interface {
	Send(*MergeRequest) error
	CloseAndRecv() (*VectorClocks, error)
	grpc.ClientStream
}

type fulcrumMergeClient struct {
	grpc.ClientStream
}

func (x *fulcrumMergeClient) Send(m *MergeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fulcrumMergeClient) CloseAndRecv() (*VectorClocks, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(VectorClocks)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fulcrumClient) MergeFulcrums(ctx context.Context, opts ...grpc.CallOption) (Fulcrum_MergeFulcrumsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Fulcrum_ServiceDesc.Streams[1], "/grpc.Fulcrum/MergeFulcrums", opts...)
	if err != nil {
		return nil, err
	}
	x := &fulcrumMergeFulcrumsClient{stream}
	return x, nil
}

type Fulcrum_MergeFulcrumsClient interface {
	Send(*MergeRequest) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type fulcrumMergeFulcrumsClient struct {
	grpc.ClientStream
}

func (x *fulcrumMergeFulcrumsClient) Send(m *MergeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fulcrumMergeFulcrumsClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FulcrumServer is the server API for Fulcrum service.
// All implementations must embed UnimplementedFulcrumServer
// for forward compatibility
type FulcrumServer interface {
	IsAvailable(context.Context, *Empty) (*Empty, error)
	// functions
	AddCity(context.Context, *AddCityRequest) (*VectorClock, error)
	DeleteCity(context.Context, *DeleteCityRequest) (*VectorClock, error)
	UpdateName(context.Context, *UpdateNameRequest) (*VectorClock, error)
	UpdateNumber(context.Context, *UpdateNumberRequest) (*VectorClock, error)
	GetNumberRebeldesFulcrum(context.Context, *GetNumberRebeldesRequest) (*GetNumberRebeldesResponse, error)
	// Merge
	VectorClockMerge(context.Context, *VectorClock) (*Empty, error)
	Merge(Fulcrum_MergeServer) error
	MergeFulcrums(Fulcrum_MergeFulcrumsServer) error
	mustEmbedUnimplementedFulcrumServer()
}

// UnimplementedFulcrumServer must be embedded to have forward compatible implementations.
type UnimplementedFulcrumServer struct {
}

func (UnimplementedFulcrumServer) IsAvailable(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAvailable not implemented")
}
func (UnimplementedFulcrumServer) AddCity(context.Context, *AddCityRequest) (*VectorClock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCity not implemented")
}
func (UnimplementedFulcrumServer) DeleteCity(context.Context, *DeleteCityRequest) (*VectorClock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCity not implemented")
}
func (UnimplementedFulcrumServer) UpdateName(context.Context, *UpdateNameRequest) (*VectorClock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateName not implemented")
}
func (UnimplementedFulcrumServer) UpdateNumber(context.Context, *UpdateNumberRequest) (*VectorClock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNumber not implemented")
}
func (UnimplementedFulcrumServer) GetNumberRebeldesFulcrum(context.Context, *GetNumberRebeldesRequest) (*GetNumberRebeldesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNumberRebeldesFulcrum not implemented")
}
func (UnimplementedFulcrumServer) VectorClockMerge(context.Context, *VectorClock) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VectorClockMerge not implemented")
}
func (UnimplementedFulcrumServer) Merge(Fulcrum_MergeServer) error {
	return status.Errorf(codes.Unimplemented, "method Merge not implemented")
}
func (UnimplementedFulcrumServer) MergeFulcrums(Fulcrum_MergeFulcrumsServer) error {
	return status.Errorf(codes.Unimplemented, "method MergeFulcrums not implemented")
}
func (UnimplementedFulcrumServer) mustEmbedUnimplementedFulcrumServer() {}

// UnsafeFulcrumServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FulcrumServer will
// result in compilation errors.
type UnsafeFulcrumServer interface {
	mustEmbedUnimplementedFulcrumServer()
}

func RegisterFulcrumServer(s grpc.ServiceRegistrar, srv FulcrumServer) {
	s.RegisterService(&Fulcrum_ServiceDesc, srv)
}

func _Fulcrum_IsAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServer).IsAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Fulcrum/IsAvailable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServer).IsAvailable(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fulcrum_AddCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServer).AddCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Fulcrum/AddCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServer).AddCity(ctx, req.(*AddCityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fulcrum_DeleteCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServer).DeleteCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Fulcrum/DeleteCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServer).DeleteCity(ctx, req.(*DeleteCityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fulcrum_UpdateName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServer).UpdateName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Fulcrum/UpdateName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServer).UpdateName(ctx, req.(*UpdateNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fulcrum_UpdateNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServer).UpdateNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Fulcrum/UpdateNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServer).UpdateNumber(ctx, req.(*UpdateNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fulcrum_GetNumberRebeldesFulcrum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNumberRebeldesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServer).GetNumberRebeldesFulcrum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Fulcrum/GetNumberRebeldesFulcrum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServer).GetNumberRebeldesFulcrum(ctx, req.(*GetNumberRebeldesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fulcrum_VectorClockMerge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VectorClock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServer).VectorClockMerge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Fulcrum/VectorClockMerge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServer).VectorClockMerge(ctx, req.(*VectorClock))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fulcrum_Merge_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FulcrumServer).Merge(&fulcrumMergeServer{stream})
}

type Fulcrum_MergeServer interface {
	SendAndClose(*VectorClocks) error
	Recv() (*MergeRequest, error)
	grpc.ServerStream
}

type fulcrumMergeServer struct {
	grpc.ServerStream
}

func (x *fulcrumMergeServer) SendAndClose(m *VectorClocks) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fulcrumMergeServer) Recv() (*MergeRequest, error) {
	m := new(MergeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Fulcrum_MergeFulcrums_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FulcrumServer).MergeFulcrums(&fulcrumMergeFulcrumsServer{stream})
}

type Fulcrum_MergeFulcrumsServer interface {
	SendAndClose(*Empty) error
	Recv() (*MergeRequest, error)
	grpc.ServerStream
}

type fulcrumMergeFulcrumsServer struct {
	grpc.ServerStream
}

func (x *fulcrumMergeFulcrumsServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fulcrumMergeFulcrumsServer) Recv() (*MergeRequest, error) {
	m := new(MergeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Fulcrum_ServiceDesc is the grpc.ServiceDesc for Fulcrum service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fulcrum_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Fulcrum",
	HandlerType: (*FulcrumServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsAvailable",
			Handler:    _Fulcrum_IsAvailable_Handler,
		},
		{
			MethodName: "AddCity",
			Handler:    _Fulcrum_AddCity_Handler,
		},
		{
			MethodName: "DeleteCity",
			Handler:    _Fulcrum_DeleteCity_Handler,
		},
		{
			MethodName: "UpdateName",
			Handler:    _Fulcrum_UpdateName_Handler,
		},
		{
			MethodName: "UpdateNumber",
			Handler:    _Fulcrum_UpdateNumber_Handler,
		},
		{
			MethodName: "GetNumberRebeldesFulcrum",
			Handler:    _Fulcrum_GetNumberRebeldesFulcrum_Handler,
		},
		{
			MethodName: "VectorClockMerge",
			Handler:    _Fulcrum_VectorClockMerge_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Merge",
			Handler:       _Fulcrum_Merge_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "MergeFulcrums",
			Handler:       _Fulcrum_MergeFulcrums_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "fulcrum.proto",
}
