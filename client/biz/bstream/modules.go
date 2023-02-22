package bstream

import (
	sem "agent/mods/interface/stream"
)

type M1 struct{}

var m1 = new(M1)

func (*M1) Change() chan bool {
	m1Chan := make(chan bool)

	//go func() {
	//	time.Sleep(10 * time.Second)
	//	m1Chan <- true
	//}()
	return m1Chan
}

func (*M1) Send(stream sem.BStreamService_SendMsgClient) error {
	return stream.Send(&sem.MsgRequest{
		Message: "M1 Hello",
	})
}
