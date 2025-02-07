// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/pbpeering/peering.proto

package pbpeering

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

// PeeringServiceClient is the client API for PeeringService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeeringServiceClient interface {
	GenerateToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*GenerateTokenResponse, error)
	Initiate(ctx context.Context, in *InitiateRequest, opts ...grpc.CallOption) (*InitiateResponse, error)
	PeeringRead(ctx context.Context, in *PeeringReadRequest, opts ...grpc.CallOption) (*PeeringReadResponse, error)
	PeeringList(ctx context.Context, in *PeeringListRequest, opts ...grpc.CallOption) (*PeeringListResponse, error)
	PeeringDelete(ctx context.Context, in *PeeringDeleteRequest, opts ...grpc.CallOption) (*PeeringDeleteResponse, error)
	// TODO(peering): As of writing, this method is only used in tests to set up Peerings in the state store.
	// Consider removing if we can find another way to populate state store in peering_endpoint_test.go
	PeeringWrite(ctx context.Context, in *PeeringWriteRequest, opts ...grpc.CallOption) (*PeeringWriteResponse, error)
	// TODO(peering): Rename this to PeeredServiceRoots? or something like that?
	TrustBundleListByService(ctx context.Context, in *TrustBundleListByServiceRequest, opts ...grpc.CallOption) (*TrustBundleListByServiceResponse, error)
	// StreamResources opens an event stream for resources to share between peers, such as services.
	// Events are streamed as they happen.
	// buf:lint:ignore RPC_REQUEST_STANDARD_NAME
	// buf:lint:ignore RPC_RESPONSE_STANDARD_NAME
	// buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE
	StreamResources(ctx context.Context, opts ...grpc.CallOption) (PeeringService_StreamResourcesClient, error)
}

type peeringServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPeeringServiceClient(cc grpc.ClientConnInterface) PeeringServiceClient {
	return &peeringServiceClient{cc}
}

func (c *peeringServiceClient) GenerateToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*GenerateTokenResponse, error) {
	out := new(GenerateTokenResponse)
	err := c.cc.Invoke(ctx, "/peering.PeeringService/GenerateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peeringServiceClient) Initiate(ctx context.Context, in *InitiateRequest, opts ...grpc.CallOption) (*InitiateResponse, error) {
	out := new(InitiateResponse)
	err := c.cc.Invoke(ctx, "/peering.PeeringService/Initiate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peeringServiceClient) PeeringRead(ctx context.Context, in *PeeringReadRequest, opts ...grpc.CallOption) (*PeeringReadResponse, error) {
	out := new(PeeringReadResponse)
	err := c.cc.Invoke(ctx, "/peering.PeeringService/PeeringRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peeringServiceClient) PeeringList(ctx context.Context, in *PeeringListRequest, opts ...grpc.CallOption) (*PeeringListResponse, error) {
	out := new(PeeringListResponse)
	err := c.cc.Invoke(ctx, "/peering.PeeringService/PeeringList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peeringServiceClient) PeeringDelete(ctx context.Context, in *PeeringDeleteRequest, opts ...grpc.CallOption) (*PeeringDeleteResponse, error) {
	out := new(PeeringDeleteResponse)
	err := c.cc.Invoke(ctx, "/peering.PeeringService/PeeringDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peeringServiceClient) PeeringWrite(ctx context.Context, in *PeeringWriteRequest, opts ...grpc.CallOption) (*PeeringWriteResponse, error) {
	out := new(PeeringWriteResponse)
	err := c.cc.Invoke(ctx, "/peering.PeeringService/PeeringWrite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peeringServiceClient) TrustBundleListByService(ctx context.Context, in *TrustBundleListByServiceRequest, opts ...grpc.CallOption) (*TrustBundleListByServiceResponse, error) {
	out := new(TrustBundleListByServiceResponse)
	err := c.cc.Invoke(ctx, "/peering.PeeringService/TrustBundleListByService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peeringServiceClient) StreamResources(ctx context.Context, opts ...grpc.CallOption) (PeeringService_StreamResourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, &PeeringService_ServiceDesc.Streams[0], "/peering.PeeringService/StreamResources", opts...)
	if err != nil {
		return nil, err
	}
	x := &peeringServiceStreamResourcesClient{stream}
	return x, nil
}

type PeeringService_StreamResourcesClient interface {
	Send(*ReplicationMessage) error
	Recv() (*ReplicationMessage, error)
	grpc.ClientStream
}

type peeringServiceStreamResourcesClient struct {
	grpc.ClientStream
}

func (x *peeringServiceStreamResourcesClient) Send(m *ReplicationMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *peeringServiceStreamResourcesClient) Recv() (*ReplicationMessage, error) {
	m := new(ReplicationMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PeeringServiceServer is the server API for PeeringService service.
// All implementations should embed UnimplementedPeeringServiceServer
// for forward compatibility
type PeeringServiceServer interface {
	GenerateToken(context.Context, *GenerateTokenRequest) (*GenerateTokenResponse, error)
	Initiate(context.Context, *InitiateRequest) (*InitiateResponse, error)
	PeeringRead(context.Context, *PeeringReadRequest) (*PeeringReadResponse, error)
	PeeringList(context.Context, *PeeringListRequest) (*PeeringListResponse, error)
	PeeringDelete(context.Context, *PeeringDeleteRequest) (*PeeringDeleteResponse, error)
	// TODO(peering): As of writing, this method is only used in tests to set up Peerings in the state store.
	// Consider removing if we can find another way to populate state store in peering_endpoint_test.go
	PeeringWrite(context.Context, *PeeringWriteRequest) (*PeeringWriteResponse, error)
	// TODO(peering): Rename this to PeeredServiceRoots? or something like that?
	TrustBundleListByService(context.Context, *TrustBundleListByServiceRequest) (*TrustBundleListByServiceResponse, error)
	// StreamResources opens an event stream for resources to share between peers, such as services.
	// Events are streamed as they happen.
	// buf:lint:ignore RPC_REQUEST_STANDARD_NAME
	// buf:lint:ignore RPC_RESPONSE_STANDARD_NAME
	// buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE
	StreamResources(PeeringService_StreamResourcesServer) error
}

// UnimplementedPeeringServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPeeringServiceServer struct {
}

func (UnimplementedPeeringServiceServer) GenerateToken(context.Context, *GenerateTokenRequest) (*GenerateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateToken not implemented")
}
func (UnimplementedPeeringServiceServer) Initiate(context.Context, *InitiateRequest) (*InitiateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Initiate not implemented")
}
func (UnimplementedPeeringServiceServer) PeeringRead(context.Context, *PeeringReadRequest) (*PeeringReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeeringRead not implemented")
}
func (UnimplementedPeeringServiceServer) PeeringList(context.Context, *PeeringListRequest) (*PeeringListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeeringList not implemented")
}
func (UnimplementedPeeringServiceServer) PeeringDelete(context.Context, *PeeringDeleteRequest) (*PeeringDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeeringDelete not implemented")
}
func (UnimplementedPeeringServiceServer) PeeringWrite(context.Context, *PeeringWriteRequest) (*PeeringWriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeeringWrite not implemented")
}
func (UnimplementedPeeringServiceServer) TrustBundleListByService(context.Context, *TrustBundleListByServiceRequest) (*TrustBundleListByServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrustBundleListByService not implemented")
}
func (UnimplementedPeeringServiceServer) StreamResources(PeeringService_StreamResourcesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamResources not implemented")
}

// UnsafePeeringServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PeeringServiceServer will
// result in compilation errors.
type UnsafePeeringServiceServer interface {
	mustEmbedUnimplementedPeeringServiceServer()
}

func RegisterPeeringServiceServer(s grpc.ServiceRegistrar, srv PeeringServiceServer) {
	s.RegisterService(&PeeringService_ServiceDesc, srv)
}

func _PeeringService_GenerateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeeringServiceServer).GenerateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peering.PeeringService/GenerateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeeringServiceServer).GenerateToken(ctx, req.(*GenerateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeeringService_Initiate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitiateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeeringServiceServer).Initiate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peering.PeeringService/Initiate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeeringServiceServer).Initiate(ctx, req.(*InitiateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeeringService_PeeringRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeeringReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeeringServiceServer).PeeringRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peering.PeeringService/PeeringRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeeringServiceServer).PeeringRead(ctx, req.(*PeeringReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeeringService_PeeringList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeeringListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeeringServiceServer).PeeringList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peering.PeeringService/PeeringList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeeringServiceServer).PeeringList(ctx, req.(*PeeringListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeeringService_PeeringDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeeringDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeeringServiceServer).PeeringDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peering.PeeringService/PeeringDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeeringServiceServer).PeeringDelete(ctx, req.(*PeeringDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeeringService_PeeringWrite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeeringWriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeeringServiceServer).PeeringWrite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peering.PeeringService/PeeringWrite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeeringServiceServer).PeeringWrite(ctx, req.(*PeeringWriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeeringService_TrustBundleListByService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrustBundleListByServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeeringServiceServer).TrustBundleListByService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peering.PeeringService/TrustBundleListByService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeeringServiceServer).TrustBundleListByService(ctx, req.(*TrustBundleListByServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeeringService_StreamResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PeeringServiceServer).StreamResources(&peeringServiceStreamResourcesServer{stream})
}

type PeeringService_StreamResourcesServer interface {
	Send(*ReplicationMessage) error
	Recv() (*ReplicationMessage, error)
	grpc.ServerStream
}

type peeringServiceStreamResourcesServer struct {
	grpc.ServerStream
}

func (x *peeringServiceStreamResourcesServer) Send(m *ReplicationMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *peeringServiceStreamResourcesServer) Recv() (*ReplicationMessage, error) {
	m := new(ReplicationMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PeeringService_ServiceDesc is the grpc.ServiceDesc for PeeringService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PeeringService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "peering.PeeringService",
	HandlerType: (*PeeringServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateToken",
			Handler:    _PeeringService_GenerateToken_Handler,
		},
		{
			MethodName: "Initiate",
			Handler:    _PeeringService_Initiate_Handler,
		},
		{
			MethodName: "PeeringRead",
			Handler:    _PeeringService_PeeringRead_Handler,
		},
		{
			MethodName: "PeeringList",
			Handler:    _PeeringService_PeeringList_Handler,
		},
		{
			MethodName: "PeeringDelete",
			Handler:    _PeeringService_PeeringDelete_Handler,
		},
		{
			MethodName: "PeeringWrite",
			Handler:    _PeeringService_PeeringWrite_Handler,
		},
		{
			MethodName: "TrustBundleListByService",
			Handler:    _PeeringService_TrustBundleListByService_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamResources",
			Handler:       _PeeringService_StreamResources_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/pbpeering/peering.proto",
}
