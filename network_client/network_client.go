package main

import (
	"fmt"

	"github.com/max_workspace/golang_learn/network"
)

func init() {
	fmt.Println("init package network_client")
}

func main() {
	// 启动simple_tcp_client
	network.SimpleTcpClient()
}
