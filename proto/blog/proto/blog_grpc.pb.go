// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: blog.proto

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

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlogServiceClient interface {
	CreatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Post, error)
	ReadPost(ctx context.Context, in *PostIdRequest, opts ...grpc.CallOption) (*Post, error)
	UpdatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Post, error)
	DeletePost(ctx context.Context, in *PostIdRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// new RPC method to retrieve all posts
	ListPosts(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (BlogService_ListPostsClient, error)
}

type blogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogServiceClient(cc grpc.ClientConnInterface) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) CreatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/blog.BlogService/CreatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ReadPost(ctx context.Context, in *PostIdRequest, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/blog.BlogService/ReadPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) UpdatePost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/blog.BlogService/UpdatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) DeletePost(ctx context.Context, in *PostIdRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/DeletePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ListPosts(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (BlogService_ListPostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &BlogService_ServiceDesc.Streams[0], "/blog.BlogService/ListPosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &blogServiceListPostsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlogService_ListPostsClient interface {
	Recv() (*PostList, error)
	grpc.ClientStream
}

type blogServiceListPostsClient struct {
	grpc.ClientStream
}

func (x *blogServiceListPostsClient) Recv() (*PostList, error) {
	m := new(PostList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlogServiceServer is the server API for BlogService service.
// All implementations must embed UnimplementedBlogServiceServer
// for forward compatibility
type BlogServiceServer interface {
	CreatePost(context.Context, *Post) (*Post, error)
	ReadPost(context.Context, *PostIdRequest) (*Post, error)
	UpdatePost(context.Context, *Post) (*Post, error)
	DeletePost(context.Context, *PostIdRequest) (*DeleteResponse, error)
	// new RPC method to retrieve all posts
	ListPosts(*EmptyRequest, BlogService_ListPostsServer) error
	mustEmbedUnimplementedBlogServiceServer()
}

// UnimplementedBlogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func (UnimplementedBlogServiceServer) CreatePost(context.Context, *Post) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedBlogServiceServer) ReadPost(context.Context, *PostIdRequest) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadPost not implemented")
}
func (UnimplementedBlogServiceServer) UpdatePost(context.Context, *Post) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (UnimplementedBlogServiceServer) DeletePost(context.Context, *PostIdRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedBlogServiceServer) ListPosts(*EmptyRequest, BlogService_ListPostsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListPosts not implemented")
}
func (UnimplementedBlogServiceServer) mustEmbedUnimplementedBlogServiceServer() {}

// UnsafeBlogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlogServiceServer will
// result in compilation errors.
type UnsafeBlogServiceServer interface {
	mustEmbedUnimplementedBlogServiceServer()
}

func RegisterBlogServiceServer(s grpc.ServiceRegistrar, srv BlogServiceServer) {
	s.RegisterService(&BlogService_ServiceDesc, srv)
}

func _BlogService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).CreatePost(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ReadPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).ReadPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/ReadPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).ReadPost(ctx, req.(*PostIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/UpdatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).UpdatePost(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/DeletePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).DeletePost(ctx, req.(*PostIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ListPosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlogServiceServer).ListPosts(m, &blogServiceListPostsServer{stream})
}

type BlogService_ListPostsServer interface {
	Send(*PostList) error
	grpc.ServerStream
}

type blogServiceListPostsServer struct {
	grpc.ServerStream
}

func (x *blogServiceListPostsServer) Send(m *PostList) error {
	return x.ServerStream.SendMsg(m)
}

// BlogService_ServiceDesc is the grpc.ServiceDesc for BlogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blog.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePost",
			Handler:    _BlogService_CreatePost_Handler,
		},
		{
			MethodName: "ReadPost",
			Handler:    _BlogService_ReadPost_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _BlogService_UpdatePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _BlogService_DeletePost_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListPosts",
			Handler:       _BlogService_ListPosts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "blog.proto",
}
