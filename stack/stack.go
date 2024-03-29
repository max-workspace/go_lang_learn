package stack

import (
	"fmt"
)

type Elementer interface{}

type stack struct {
	element []Elementer
}

func init() {
	fmt.Println("init stack")
}

func NewStack() *stack {
	return &stack{}
}

func (this *stack) Size() int {
	return len(this.element)
}

func (this *stack) Push(e Elementer) {
	this.element = append(this.element, e)
}

func (this *stack) Pop() Elementer {
	size := this.Size()
	if size == 0 {
		fmt.Println("stack is empty!")
		return nil
	}
	lastElement := this.element[size-1]
	this.element[size-1] = nil
	this.element = this.element[:size-1]
	return lastElement
}

func (this *stack) IsEmpty() bool {
	if len(this.element) == 0 {
		return true
	}
	return false
}

func (this *stack) Clear() bool {
	if this.IsEmpty() {
		fmt.Println("stack is empty!")
		return false
	}
	for i := 0; i < this.Size(); i++ {
		this.element[i] = nil
	}
	this.element = make([]Elementer, 0)
	return true
}
