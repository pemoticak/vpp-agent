// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package generic

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MetaServiceClient is the client API for MetaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetaServiceClient interface {
	// KnownModels returns information about service capabilities
	// including list of models supported by the server.
	KnownModels(ctx context.Context, in *KnownModelsRequest, opts ...grpc.CallOption) (*KnownModelsResponse, error)
}

type metaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetaServiceClient(cc grpc.ClientConnInterface) MetaServiceClient {
	return &metaServiceClient{cc}
}

func (c *metaServiceClient) KnownModels(ctx context.Context, in *KnownModelsRequest, opts ...grpc.CallOption) (*KnownModelsResponse, error) {
	out := new(KnownModelsResponse)
	err := c.cc.Invoke(ctx, "/ligato.generic.MetaService/KnownModels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetaServiceServer is the server API for MetaService service.
// All implementations must embed UnimplementedMetaServiceServer
// for forward compatibility
type MetaServiceServer interface {
	// KnownModels returns information about service capabilities
	// including list of models supported by the server.
	KnownModels(context.Context, *KnownModelsRequest) (*KnownModelsResponse, error)
	mustEmbedUnimplementedMetaServiceServer()
}

// UnimplementedMetaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMetaServiceServer struct {
}

func (*UnimplementedMetaServiceServer) KnownModels(context.Context, *KnownModelsRequest) (*KnownModelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KnownModels not implemented")
}
func (*UnimplementedMetaServiceServer) mustEmbedUnimplementedMetaServiceServer() {}

func RegisterMetaServiceServer(s *grpc.Server, srv MetaServiceServer) {
	s.RegisterService(&_MetaService_serviceDesc, srv)
}

func _MetaService_KnownModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KnownModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).KnownModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ligato.generic.MetaService/KnownModels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).KnownModels(ctx, req.(*KnownModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ligato.generic.MetaService",
	HandlerType: (*MetaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "KnownModels",
			Handler:    _MetaService_KnownModels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ligato/generic/meta.proto",
}