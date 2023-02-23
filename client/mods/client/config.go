package client

import "agent/mods/config"

const _key = "server"

type ConnectServerConfig struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}

func regConfig() {
	c := ConnectServerConfig{
		Ip:   "127.0.0.1",
		Port: 9088,
	}
	_ = config.RegConfig(_key, c)
}

func GetConnectServerConfig() ConnectServerConfig {
	var c ConnectServerConfig
	_ = config.GetConfig(_key, &c)
	return c
}

func init() {
	regConfig()
}
