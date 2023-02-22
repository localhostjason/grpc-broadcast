package node

import (
	"console/mods/interface/node"
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type NodeInfo struct {
	node.UnimplementedNodeInfoServiceServer
}

// GetNodeInfo response
func (n *NodeInfo) GetNodeInfo(ctx context.Context, in *node.NodeRequest) (*node.NodeResponse, error) {
	log.Printf("Receive message body from client: %s", in)

	return &node.NodeResponse{Code: "SUCCESS", Message: "ok"}, nil
}

func RegisterNodeService(gs *grpc.Server) {
	node.RegisterNodeInfoServiceServer(gs, &NodeInfo{})
}
