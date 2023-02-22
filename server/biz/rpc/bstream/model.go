package bstream

import (
	"console/mods/interface/stream"
	"github.com/golang/protobuf/ptypes/timestamp"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

// ConnectPool 定义一个类，继承字典(异步,带锁的),一会存入grpc stream对象 { name : stream<obj> }
type ConnectPool struct {
	sync.Map
}

// Get 为这个 类<对象池> 添加方法，分别为 Get,Add,Del和BroadCast(广播信息,群发)
func (p *ConnectPool) Get(name string) stream.BStreamService_SendMsgServer {
	if s, ok := p.Load(name); ok {
		return s.(stream.BStreamService_SendMsgServer)
	} else {
		return nil
	}
}

func (p *ConnectPool) Add(name string, stream stream.BStreamService_SendMsgServer) {
	p.Store(name, stream)
}

func (p *ConnectPool) Del(name string) {
	p.Delete(name)
}

// BroadCast  广播 发送消息, 双向流 实时交互
func (p *ConnectPool) BroadCast(from, message string) {
	log.Printf("BroadCast from: %s, message: %s\n", from, message)
	p.Range(func(usernameI, streamI interface{}) bool {
		username := usernameI.(string)
		s := streamI.(stream.BStreamService_SendMsgServer)
		if username != from {
			return true
		} else {
			_ = s.Send(&stream.MsgReply{
				Message:     message,
				MessageType: stream.MsgReply_NORMAL_MESSAGE, // 2.正常数据
				TS:          &timestamp.Timestamp{Seconds: time.Now().Unix()},
			})
		}
		return true
	})
}
