// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package stream

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BStreamServiceClient is the client API for BStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BStreamServiceClient interface {
	SendMsg(ctx context.Context, opts ...grpc.CallOption) (BStreamService_SendMsgClient, error)
}

type bStreamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBStreamServiceClient(cc grpc.ClientConnInterface) BStreamServiceClient {
	return &bStreamServiceClient{cc}
}

func (c *bStreamServiceClient) SendMsg(ctx context.Context, opts ...grpc.CallOption) (BStreamService_SendMsgClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BStreamService_serviceDesc.Streams[0], "/bstream.BStreamService/SendMsg", opts...)
	if err != nil {
		return nil, err
	}
	x := &bStreamServiceSendMsgClient{stream}
	return x, nil
}

type BStreamService_SendMsgClient interface {
	Send(*MsgRequest) error
	Recv() (*MsgReply, error)
	grpc.ClientStream
}

type bStreamServiceSendMsgClient struct {
	grpc.ClientStream
}

func (x *bStreamServiceSendMsgClient) Send(m *MsgRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bStreamServiceSendMsgClient) Recv() (*MsgReply, error) {
	m := new(MsgReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BStreamServiceServer is the server API for BStreamService service.
// All implementations must embed UnimplementedBStreamServiceServer
// for forward compatibility
type BStreamServiceServer interface {
	SendMsg(BStreamService_SendMsgServer) error
	mustEmbedUnimplementedBStreamServiceServer()
}

// UnimplementedBStreamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBStreamServiceServer struct {
}

func (*UnimplementedBStreamServiceServer) SendMsg(BStreamService_SendMsgServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMsg not implemented")
}
func (*UnimplementedBStreamServiceServer) mustEmbedUnimplementedBStreamServiceServer() {}

func RegisterBStreamServiceServer(s *grpc.Server, srv BStreamServiceServer) {
	s.RegisterService(&_BStreamService_serviceDesc, srv)
}

func _BStreamService_SendMsg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BStreamServiceServer).SendMsg(&bStreamServiceSendMsgServer{stream})
}

type BStreamService_SendMsgServer interface {
	Send(*MsgReply) error
	Recv() (*MsgRequest, error)
	grpc.ServerStream
}

type bStreamServiceSendMsgServer struct {
	grpc.ServerStream
}

func (x *bStreamServiceSendMsgServer) Send(m *MsgReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bStreamServiceSendMsgServer) Recv() (*MsgRequest, error) {
	m := new(MsgRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _BStreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bstream.BStreamService",
	HandlerType: (*BStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMsg",
			Handler:       _BStreamService_SendMsg_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "bstream.proto",
}