package client

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var GCONN *grpc.ClientConn

type Client struct {
}

func (c *Client) Connect() error {
	cfg := GetConnectServerConfig()
	addr := fmt.Sprintf("%s:%d", cfg.Ip, cfg.Port)

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return conn.Close()
	}

	GCONN = conn
	return nil
}

func closeConn() error {
	if GCONN == nil {
		return nil
	}
	return GCONN.Close()
}

func Connect() error {
	conn := &Client{}
	return conn.Connect()
}

func Close() {
	_ = closeConn()
}
