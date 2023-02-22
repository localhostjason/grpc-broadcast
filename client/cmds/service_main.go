package cmds

import (
	"agent/biz/bstream"
	"agent/mods/client"
	"agent/mods/db"
	"agent/mods/svc"
	"errors"
	"os"
	"syscall"

	log "github.com/sirupsen/logrus"
)

type MainProc struct {
	singleMode bool
	quit       chan bool
}

func (m *MainProc) Stop() {
	m.quit <- true
}

func NewProc(singleMode bool) *MainProc {
	return &MainProc{singleMode: singleMode, quit: make(chan bool)}
}

func (m *MainProc) Run(svc *svc.Svc) {
	if err := startAgent(); err != nil {
		return
	}

	cbs := bstream.NewCheckBStreamCon()
	cbs.Check()

	<-m.quit
	cbs.Stop()

	client.Close()
}

func (m *MainProc) SigHandlers() map[os.Signal]svc.SignalHandlerFunc {
	return map[os.Signal]svc.SignalHandlerFunc{
		syscall.SIGTERM: m.handleSigTerm,
		os.Interrupt:    m.handleSigTerm,
	}
}

func (m *MainProc) handleSigTerm(sig os.Signal) (err error) {
	m.quit <- true
	return errors.New("quit by signal " + sig.String())
}

func startAgent() error {
	if err := db.Connect(); err != nil {
		log.Fatalln(err)
	}
	if err := db.InitData(); err != nil {
		log.Fatalln(err)
	}

	if err := client.Connect(); err != nil {
		log.Fatalln(err)
	}
	return nil
}
