// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0--rc2
// source: product.proto

package product

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

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductClient interface {
	Products(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
}

type productClient struct {
	cc grpc.ClientConnInterface
}

func NewProductClient(cc grpc.ClientConnInterface) ProductClient {
	return &productClient{cc}
}

func (c *productClient) Products(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/product.Product/Products", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
// All implementations must embed UnimplementedProductServer
// for forward compatibility
type ProductServer interface {
	Products(context.Context, *ProductRequest) (*ProductResponse, error)
	mustEmbedUnimplementedProductServer()
}

// UnimplementedProductServer must be embedded to have forward compatible implementations.
type UnimplementedProductServer struct {
}

func (UnimplementedProductServer) Products(context.Context, *ProductRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Products not implemented")
}
func (UnimplementedProductServer) mustEmbedUnimplementedProductServer() {}

// UnsafeProductServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServer will
// result in compilation errors.
type UnsafeProductServer interface {
	mustEmbedUnimplementedProductServer()
}

func RegisterProductServer(s grpc.ServiceRegistrar, srv ProductServer) {
	s.RegisterService(&Product_ServiceDesc, srv)
}

func _Product_Products_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).Products(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.Product/Products",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).Products(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Product_ServiceDesc is the grpc.ServiceDesc for Product service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Product_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.Product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Products",
			Handler:    _Product_Products_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
