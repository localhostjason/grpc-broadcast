package bstream

import (
	"agent/mods/client"
	sem "agent/mods/interface/stream"
	"context"
	"errors"
	"fmt"
	"io"

	//uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

type BStream struct {
}

func NewBStream() *BStream {
	return &BStream{}
}

func (c *BStream) Work(stream sem.BStreamService_SendMsgClient, cancel context.CancelFunc) {
	// 创建了一个连接管道
	connected := make(chan bool)

	// 接收 服务端信息
	go c.replyMsg(stream, connected, cancel)

	// 发送消息
	go c.sendMsg(stream, connected)
}

func (c *BStream) createStream() (sem.BStreamService_SendMsgClient, context.CancelFunc, error) {
	newSC := sem.NewBStreamServiceClient(client.GCONN)
	// 声明 context
	ctx, cancel := context.WithCancel(context.Background())
	// name set = uuid
	//clientId := uuid.NewV4().String()
	clientId := "UUID-01"
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("uid", clientId))

	// 创建双向数据流
	stream, err := newSC.SendMsg(ctx)

	if err != nil {
		// err grpc error: code = Unavailable desc = connection closed before server preface received ?
		// time sleep 3 second
		cancel()
		return stream, cancel, errors.New(fmt.Sprintf("创建数据流失败: [%v] ", err))
	}
	return stream, cancel, nil
}

// 收到服务器消息， 收到消息后 去做某些事情
func (c *BStream) replyMsg(stream sem.BStreamService_SendMsgClient, connected chan bool, cancel context.CancelFunc) {
	var (
		reply *sem.MsgReply
		err   error
	)
	for {
		reply, err = stream.Recv()
		if err != nil {
			log.Debug("reply msg err:", err)
			cancel()
			ActiveServer = false
			break
		}
		log.Info("reply msg: ", reply.Message)

		if reply.MessageType == sem.MsgReply_CONNECT_FAILED { // code=1 连接失败
			cancel()
			ActiveServer = false
			break
		}
		if reply.MessageType == sem.MsgReply_CONNECT_SUCCESS { // code=0 连接成功
			connected <- true
		}
		// 基本都是两个if都不执行，去下一次循环,返回的是 code=2 正常消息
	}
}

// 发送消息 (client 发生消息 给服务： 暂时不需要，让客户端 去发送请求 给 服务器。)
func (c *BStream) sendMsg(stream sem.BStreamService_SendMsgClient, connected chan bool) {
	<-connected

	for {
		select {
		case <-m1.Change():
			err := m1.Send(stream)
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Error("send msg err:", err)
			}
		}
	}

}

func RunBStream() {
	//  err grpc error: code = Unavailable desc = connection closed before server preface received ?
	// time.sleep 3 second
}
