// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: stroeer/navigation/v1/navigation_service.proto

package navigation

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

const (
	NavigationService_GetNavigation_FullMethodName = "/stroeer.navigation.v1.NavigationService/GetNavigation"
)

// NavigationServiceClient is the client API for NavigationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NavigationServiceClient interface {
	// Retrieve navigation tree structure
	GetNavigation(ctx context.Context, in *GetNavigationRequest, opts ...grpc.CallOption) (*GetNavigationResponse, error)
}

type navigationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNavigationServiceClient(cc grpc.ClientConnInterface) NavigationServiceClient {
	return &navigationServiceClient{cc}
}

func (c *navigationServiceClient) GetNavigation(ctx context.Context, in *GetNavigationRequest, opts ...grpc.CallOption) (*GetNavigationResponse, error) {
	out := new(GetNavigationResponse)
	err := c.cc.Invoke(ctx, NavigationService_GetNavigation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NavigationServiceServer is the server API for NavigationService service.
// All implementations must embed UnimplementedNavigationServiceServer
// for forward compatibility
type NavigationServiceServer interface {
	// Retrieve navigation tree structure
	GetNavigation(context.Context, *GetNavigationRequest) (*GetNavigationResponse, error)
	mustEmbedUnimplementedNavigationServiceServer()
}

// UnimplementedNavigationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNavigationServiceServer struct {
}

func (UnimplementedNavigationServiceServer) GetNavigation(context.Context, *GetNavigationRequest) (*GetNavigationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNavigation not implemented")
}
func (UnimplementedNavigationServiceServer) mustEmbedUnimplementedNavigationServiceServer() {}

// UnsafeNavigationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NavigationServiceServer will
// result in compilation errors.
type UnsafeNavigationServiceServer interface {
	mustEmbedUnimplementedNavigationServiceServer()
}

func RegisterNavigationServiceServer(s grpc.ServiceRegistrar, srv NavigationServiceServer) {
	s.RegisterService(&NavigationService_ServiceDesc, srv)
}

func _NavigationService_GetNavigation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNavigationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NavigationServiceServer).GetNavigation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NavigationService_GetNavigation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NavigationServiceServer).GetNavigation(ctx, req.(*GetNavigationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NavigationService_ServiceDesc is the grpc.ServiceDesc for NavigationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NavigationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stroeer.navigation.v1.NavigationService",
	HandlerType: (*NavigationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNavigation",
			Handler:    _NavigationService_GetNavigation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stroeer/navigation/v1/navigation_service.proto",
}