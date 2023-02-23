package cmds

import (
	"console/biz/rpc/bstream"
	"console/biz/rpc/node"
	"console/biz/view"
	"google.golang.org/grpc"
)

func Run() {
	s := NewMainServer()
	s.LoadView(view.SetView)

	s.LoadGrpcServerApi(func(server *grpc.Server) {
		bstream.RegisterBStreamService(server)
		node.RegisterNodeService(server)
	})

	s.Run()
}
