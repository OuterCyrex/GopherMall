// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: fav.proto

package __FavProto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Fav_GetFavList_FullMethodName       = "/fav.Fav/GetFavList"
	Fav_AddUserFav_FullMethodName       = "/fav.Fav/AddUserFav"
	Fav_DeleteUserFav_FullMethodName    = "/fav.Fav/DeleteUserFav"
	Fav_GetUserFavDetail_FullMethodName = "/fav.Fav/GetUserFavDetail"
)

// FavClient is the client API for Fav service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FavClient interface {
	GetFavList(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*UserFavListResponse, error)
	AddUserFav(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*Empty, error)
	DeleteUserFav(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*Empty, error)
	GetUserFavDetail(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*Empty, error)
}

type favClient struct {
	cc grpc.ClientConnInterface
}

func NewFavClient(cc grpc.ClientConnInterface) FavClient {
	return &favClient{cc}
}

func (c *favClient) GetFavList(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*UserFavListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserFavListResponse)
	err := c.cc.Invoke(ctx, Fav_GetFavList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favClient) AddUserFav(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Fav_AddUserFav_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favClient) DeleteUserFav(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Fav_DeleteUserFav_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favClient) GetUserFavDetail(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Fav_GetUserFavDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FavServer is the server API for Fav service.
// All implementations must embed UnimplementedFavServer
// for forward compatibility.
type FavServer interface {
	GetFavList(context.Context, *UserFavRequest) (*UserFavListResponse, error)
	AddUserFav(context.Context, *UserFavRequest) (*Empty, error)
	DeleteUserFav(context.Context, *UserFavRequest) (*Empty, error)
	GetUserFavDetail(context.Context, *UserFavRequest) (*Empty, error)
	mustEmbedUnimplementedFavServer()
}

// UnimplementedFavServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFavServer struct{}

func (UnimplementedFavServer) GetFavList(context.Context, *UserFavRequest) (*UserFavListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavList not implemented")
}
func (UnimplementedFavServer) AddUserFav(context.Context, *UserFavRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserFav not implemented")
}
func (UnimplementedFavServer) DeleteUserFav(context.Context, *UserFavRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserFav not implemented")
}
func (UnimplementedFavServer) GetUserFavDetail(context.Context, *UserFavRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFavDetail not implemented")
}
func (UnimplementedFavServer) mustEmbedUnimplementedFavServer() {}
func (UnimplementedFavServer) testEmbeddedByValue()             {}

// UnsafeFavServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FavServer will
// result in compilation errors.
type UnsafeFavServer interface {
	mustEmbedUnimplementedFavServer()
}

func RegisterFavServer(s grpc.ServiceRegistrar, srv FavServer) {
	// If the following call pancis, it indicates UnimplementedFavServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Fav_ServiceDesc, srv)
}

func _Fav_GetFavList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFavRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavServer).GetFavList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Fav_GetFavList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavServer).GetFavList(ctx, req.(*UserFavRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fav_AddUserFav_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFavRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavServer).AddUserFav(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Fav_AddUserFav_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavServer).AddUserFav(ctx, req.(*UserFavRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fav_DeleteUserFav_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFavRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavServer).DeleteUserFav(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Fav_DeleteUserFav_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavServer).DeleteUserFav(ctx, req.(*UserFavRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fav_GetUserFavDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFavRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavServer).GetUserFavDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Fav_GetUserFavDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavServer).GetUserFavDetail(ctx, req.(*UserFavRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Fav_ServiceDesc is the grpc.ServiceDesc for Fav service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fav_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fav.Fav",
	HandlerType: (*FavServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFavList",
			Handler:    _Fav_GetFavList_Handler,
		},
		{
			MethodName: "AddUserFav",
			Handler:    _Fav_AddUserFav_Handler,
		},
		{
			MethodName: "DeleteUserFav",
			Handler:    _Fav_DeleteUserFav_Handler,
		},
		{
			MethodName: "GetUserFavDetail",
			Handler:    _Fav_GetUserFavDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fav.proto",
}
