package bstream

import (
	sem "console/mods/interface/stream"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

var ConnectPools = new(ConnectPool)

// Service 定义服务器类
type Service struct {
	sem.UnimplementedBStreamServiceServer
	Uid string
}

func (s *Service) Connect(stream sem.BStreamService_SendMsgServer) error {
	peerCtx, _ := peer.FromContext(stream.Context())
	log.Printf("Received new connection.  %s", peerCtx.Addr.String())

	md, _ := metadata.FromIncomingContext(stream.Context())
	fmt.Println("meta data:", md)
	uid := md["uid"][0] // 从metadata中获取客户端身份信息，可以理解为请求头里的数据
	s.Uid = uid

	if ConnectPools.Get(uid) != nil {
		return stream.Send(&sem.MsgReply{
			Message:     fmt.Sprintf("username %s already exists!", uid),
			MessageType: sem.MsgReply_CONNECT_FAILED, // 1. 连接失败 ， 重名了 用户已经存在
		})

	}

	ConnectPools.Add(uid, stream)
	return stream.Send(&sem.MsgReply{
		Message:     fmt.Sprintf("Connect success!"),
		MessageType: sem.MsgReply_CONNECT_SUCCESS, // 0 连接成功
	})

}

func (s *Service) SendMsg(stream sem.BStreamService_SendMsgServer) error {
	if err := s.Connect(stream); err != nil {
		return err
	}

	go func() {
		// 阻塞住，等待断开连接的时候触发
		<-stream.Context().Done()
		ConnectPools.Del(s.Uid)
		ConnectPools.BroadCast(s.Uid, fmt.Sprintf("%s leval room", s.Uid))
	}()

	// 广播
	ConnectPools.BroadCast(s.Uid, fmt.Sprintf("Welcome %s!", s.Uid))

	//  阻塞接收 该用户后续传来的消息
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Info("req msg: ", req.Message)
	}
}

func RegisterBStreamService(gs *grpc.Server) {
	sem.RegisterBStreamServiceServer(gs, &Service{})
}
