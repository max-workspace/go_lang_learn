package example

import (
	"fmt"
)

func init() {
	fmt.Println("init example counter")
}

func TestCounter() {
	fmt.Println("test counter")
	ch := make(chan int)
	go GetCounter(ch)
	ShowCounter(ch)
}

func GetCounter(ch chan int) {
	for i := 0; i < 3; i++ {
		ch <- i
	}
	defer close(ch)
}

func ShowCounter(ch chan int) {
	for {
		if i, ok := <-ch; ok {
			fmt.Printf("ok is %t and the counter is at %d\n", ok, i)
		} else {
			break
		}
	}
}
