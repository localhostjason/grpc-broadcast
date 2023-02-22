package cmds

import (
	"console/biz/rpc/bstream"
	"console/biz/rpc/node"
	"console/biz/view"
	"console/mods/casbinx"
	"github.com/localhostjason/webserver/daemonx"
	"google.golang.org/grpc"
)

func Run() {
	s := daemonx.NewMainServer()

	s.LoadPluginHandler(casbinx.NewCasBin().Run)
	s.LoadView(view.SetView)

	s.LoadGrpcServerApi(func(server *grpc.Server) {
		bstream.RegisterBStreamService(server)
		node.RegisterNodeService(server)
	})

	s.Run()
}
