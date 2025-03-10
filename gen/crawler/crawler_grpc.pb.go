// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: crawler/crawler.proto

package crawler

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
	CrawlerService_StartCrawl_FullMethodName   = "/crawler.CrawlerService/StartCrawl"
	CrawlerService_GetJobStatus_FullMethodName = "/crawler.CrawlerService/GetJobStatus"
)

// CrawlerServiceClient is the client API for CrawlerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CrawlerServiceClient interface {
	StartCrawl(ctx context.Context, in *CrawlRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[CrawlResponse], error)
	GetJobStatus(ctx context.Context, in *JobStatusRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[JobStatusResponse], error)
}

type crawlerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCrawlerServiceClient(cc grpc.ClientConnInterface) CrawlerServiceClient {
	return &crawlerServiceClient{cc}
}

func (c *crawlerServiceClient) StartCrawl(ctx context.Context, in *CrawlRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[CrawlResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CrawlerService_ServiceDesc.Streams[0], CrawlerService_StartCrawl_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[CrawlRequest, CrawlResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CrawlerService_StartCrawlClient = grpc.ServerStreamingClient[CrawlResponse]

func (c *crawlerServiceClient) GetJobStatus(ctx context.Context, in *JobStatusRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[JobStatusResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CrawlerService_ServiceDesc.Streams[1], CrawlerService_GetJobStatus_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[JobStatusRequest, JobStatusResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CrawlerService_GetJobStatusClient = grpc.ServerStreamingClient[JobStatusResponse]

// CrawlerServiceServer is the server API for CrawlerService service.
// All implementations must embed UnimplementedCrawlerServiceServer
// for forward compatibility.
type CrawlerServiceServer interface {
	StartCrawl(*CrawlRequest, grpc.ServerStreamingServer[CrawlResponse]) error
	GetJobStatus(*JobStatusRequest, grpc.ServerStreamingServer[JobStatusResponse]) error
	mustEmbedUnimplementedCrawlerServiceServer()
}

// UnimplementedCrawlerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCrawlerServiceServer struct{}

func (UnimplementedCrawlerServiceServer) StartCrawl(*CrawlRequest, grpc.ServerStreamingServer[CrawlResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StartCrawl not implemented")
}
func (UnimplementedCrawlerServiceServer) GetJobStatus(*JobStatusRequest, grpc.ServerStreamingServer[JobStatusResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetJobStatus not implemented")
}
func (UnimplementedCrawlerServiceServer) mustEmbedUnimplementedCrawlerServiceServer() {}
func (UnimplementedCrawlerServiceServer) testEmbeddedByValue()                        {}

// UnsafeCrawlerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CrawlerServiceServer will
// result in compilation errors.
type UnsafeCrawlerServiceServer interface {
	mustEmbedUnimplementedCrawlerServiceServer()
}

func RegisterCrawlerServiceServer(s grpc.ServiceRegistrar, srv CrawlerServiceServer) {
	// If the following call pancis, it indicates UnimplementedCrawlerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CrawlerService_ServiceDesc, srv)
}

func _CrawlerService_StartCrawl_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CrawlRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CrawlerServiceServer).StartCrawl(m, &grpc.GenericServerStream[CrawlRequest, CrawlResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CrawlerService_StartCrawlServer = grpc.ServerStreamingServer[CrawlResponse]

func _CrawlerService_GetJobStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JobStatusRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CrawlerServiceServer).GetJobStatus(m, &grpc.GenericServerStream[JobStatusRequest, JobStatusResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CrawlerService_GetJobStatusServer = grpc.ServerStreamingServer[JobStatusResponse]

// CrawlerService_ServiceDesc is the grpc.ServiceDesc for CrawlerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CrawlerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crawler.CrawlerService",
	HandlerType: (*CrawlerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartCrawl",
			Handler:       _CrawlerService_StartCrawl_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetJobStatus",
			Handler:       _CrawlerService_GetJobStatus_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "crawler/crawler.proto",
}
