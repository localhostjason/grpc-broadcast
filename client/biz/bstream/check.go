package bstream

import (
	log "github.com/sirupsen/logrus"
	"time"
)

var ActiveServer = false

type CheckBStreamCon struct {
	Quit chan bool
}

func NewCheckBStreamCon() *CheckBStreamCon {
	return &CheckBStreamCon{
		Quit: make(chan bool),
	}
}

func (c *CheckBStreamCon) Check() {
	go c.work()
}

func (c *CheckBStreamCon) work() {
	ticker := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-ticker.C:
			if ActiveServer {
				continue
			}
			log.Info("check client connect")
			bs := NewBStream()
			stream, cancel, err := bs.createStream()
			if err != nil {
				log.Error(err)
				ActiveServer = false
			} else {
				bs.Work(stream, cancel)
				ActiveServer = true
			}
		case <-c.Quit:
			ticker.Stop()
			log.Info("stop check connect")
		}
	}
}

func (c *CheckBStreamCon) Stop() {
	c.Quit <- true
}
