package example

import (
	"fmt"
)

func init() {
	fmt.Println("init example range")
}

func TestRange() {
	fmt.Println("test range")
	ch := make(chan int)
	go RangeGeneration(ch)
	ShowRangeGeneration(ch)
}

func RangeGeneration(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	defer close(ch)
}

func ShowRangeGeneration(ch chan int) {
	for v := range ch {
		fmt.Printf("The value is %v\n", v)
	}
	fmt.Println("ShowRangeGeneration finish")
}
