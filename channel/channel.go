package main

import (
	"fmt"

	"github.com/max_workspace/golang_learn/channel/example"
)

func init() {
	fmt.Println("init package channel")
}

func main() {
	// 测试通道的range
	example.TestRange()

	// 测试select
	example.TestSelect()

	// 测试通道的阻塞
	example.TestCounter()

}
