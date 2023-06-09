// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: bookstore/bookstore.proto

package bookstore

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

// BookStoreServiceClient is the client API for BookStoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookStoreServiceClient interface {
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*CreateBookResponse, error)
	GetBook(ctx context.Context, in *GetBookListRequest, opts ...grpc.CallOption) (*GetBookListResponse, error)
}

type bookStoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookStoreServiceClient(cc grpc.ClientConnInterface) BookStoreServiceClient {
	return &bookStoreServiceClient{cc}
}

func (c *bookStoreServiceClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*CreateBookResponse, error) {
	out := new(CreateBookResponse)
	err := c.cc.Invoke(ctx, "/bookstore.BookStoreService/CreateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookStoreServiceClient) GetBook(ctx context.Context, in *GetBookListRequest, opts ...grpc.CallOption) (*GetBookListResponse, error) {
	out := new(GetBookListResponse)
	err := c.cc.Invoke(ctx, "/bookstore.BookStoreService/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookStoreServiceServer is the server API for BookStoreService service.
// All implementations must embed UnimplementedBookStoreServiceServer
// for forward compatibility
type BookStoreServiceServer interface {
	CreateBook(context.Context, *CreateBookRequest) (*CreateBookResponse, error)
	GetBook(context.Context, *GetBookListRequest) (*GetBookListResponse, error)
	mustEmbedUnimplementedBookStoreServiceServer()
}

// UnimplementedBookStoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookStoreServiceServer struct {
}

func (UnimplementedBookStoreServiceServer) CreateBook(context.Context, *CreateBookRequest) (*CreateBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (UnimplementedBookStoreServiceServer) GetBook(context.Context, *GetBookListRequest) (*GetBookListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedBookStoreServiceServer) mustEmbedUnimplementedBookStoreServiceServer() {}

// UnsafeBookStoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookStoreServiceServer will
// result in compilation errors.
type UnsafeBookStoreServiceServer interface {
	mustEmbedUnimplementedBookStoreServiceServer()
}

func RegisterBookStoreServiceServer(s grpc.ServiceRegistrar, srv BookStoreServiceServer) {
	s.RegisterService(&BookStoreService_ServiceDesc, srv)
}

func _BookStoreService_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookStoreServiceServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.BookStoreService/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookStoreServiceServer).CreateBook(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookStoreService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookStoreServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.BookStoreService/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookStoreServiceServer).GetBook(ctx, req.(*GetBookListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookStoreService_ServiceDesc is the grpc.ServiceDesc for BookStoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookStoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bookstore.BookStoreService",
	HandlerType: (*BookStoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBook",
			Handler:    _BookStoreService_CreateBook_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _BookStoreService_GetBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bookstore/bookstore.proto",
}
