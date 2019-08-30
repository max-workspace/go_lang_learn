package main

import (
	"fmt"

	"github.com/max_workspace/golang_learn/stack"
	"github.com/max_workspace/golang_learn/struct/people"
)

type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

func main() {
	// 测试外部导入的struct
	testStructPeople()

	// 测试嵌套的struct
	testStructVoodoo()

	// 测试stack
	testStructStack()
}

func testStructPeople() {
	fmt.Println("test people struct")
	max := people.NewPeople("max", 26, 1)
	fmt.Printf(
		"type:%T|name:%v|age:%d|gender:%v|createTime:%v|\n",
		max,
		max.Name,
		max.Age(),
		max.GenderString(),
		max.CreateTime())
	fmt.Println("\n")
}

func testStructVoodoo() {
	fmt.Println("test Voodoo struct")
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
	fmt.Println("\n")
}

func testStructStack() {
	fmt.Println("test stack struct")
	entry := stack.NewStack()
	for i := 0; i < 50; i++ {
		entry.Push(i)
	}
	fmt.Println(entry.Size())
	fmt.Println(entry.Pop())
	fmt.Println(entry.Pop(), entry.Size())
	fmt.Println(entry.Clear())
	for i := 0; i < 50; i++ {
		entry.Push(i)
	}
	fmt.Println(entry.Size())

}
