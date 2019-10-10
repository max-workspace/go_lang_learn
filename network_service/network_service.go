package main

import (
	"fmt"

	"github.com/max_workspace/golang_learn/network"
)

func init() {
	fmt.Println("init package network_service")
}

func main() {
	// 启动simple_tcp_service
	// network.SimpleTcpService()

	// 启动simple_web_service
	// network.SimpleWebService()

	// 启动多功能网页服务器
	network.ElaboratedWebserver()
}
