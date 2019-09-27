package network

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func init() {
	fmt.Println("init network simple_tcp_client")
}

func SimpleTcpClient() {
	// 打开连接
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		// 由于目标计算机拒绝而无法创建连接
		fmt.Println("Error dialing", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\r\n")
	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
	}
}
