package node

import (
	"agent/mods/client"
	"agent/mods/interface/node"
	"context"
	"fmt"
)

func LoadNode() {

	c := node.NewNodeInfoServiceClient(client.GCONN)

	response, err := c.GetNodeInfo(context.Background(), &node.NodeRequest{
		NodeId: "ABC",
	})

	if err != nil {
		fmt.Printf("Error when calling getNodeInfo: %s\n", err)
		return
	}
	fmt.Println("response ===> ", response)
}
