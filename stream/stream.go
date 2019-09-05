package main

import (
	"fmt"

	"github.com/max_workspace/golang_learn/stream/file"
	"github.com/max_workspace/golang_learn/stream/input"
)

func main() {
	fmt.Println("package stream")

	// 测试从用户获取输入
	// input.TestInputFromUser()

	// 测试从缓存中获取输入
	// input.TestInputFromBuffer()

	// 测试从字符串获取输入
	input.TestInputFromString()

	// 测试从文件中读取
	file.TestReadFromFile()
}
