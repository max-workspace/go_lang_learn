package main

import (
	"fmt"
	"time"

	"github.com/max_workspace/golang_learn/stack"
	"github.com/max_workspace/golang_learn/struct/people"
)

// 通过声明接口，实现定义未导入类型的map
type Peopler interface {
	Name() string
	Age() int
	GenderString() string
	CreateTime() time.Time
}

type Base struct{}

type Voodoo struct {
	Base
}

var (
	PeopleList = make(map[string]Peopler)
)

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
	createPeople("max", 26, 1)
	createPeople("sum", 27, 1)
	fmt.Printf("%v", PeopleList)
	showPeopleInfo("max")
	showPeopleInfo("sum")
	fmt.Println("\n")
}

func createPeople(name string, age int, gender int) {
	PeopleList[name] = people.NewPeople(name, age, gender)
}

func showPeopleInfo(name string) {
	fmt.Printf(
		"%v|type:%T|name:%v|age:%d|gender:%v|createTime:%v|\n",
		name,
		PeopleList[name],
		PeopleList[name].Name(),
		PeopleList[name].Age(),
		PeopleList[name].GenderString(),
		PeopleList[name].CreateTime())
}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
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
