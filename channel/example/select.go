package example

import (
	"fmt"
)

func init() {
	fmt.Println("init example select")
}

func TestSelect() {
	fmt.Println("test select")
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	suck(ch1, ch2)
}

func pump1(ch chan int) {
	for i := 0; i < 3; i++ {
		ch <- i * 2
	}
	defer close(ch)
}

func pump2(ch chan int) {
	for i := 0; i < 3; i++ {
		ch <- i * 3
	}
	defer close(ch)
}

func suck(ch1, ch2 chan int) {
	for i := 0; i < 6; i++ {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		}
	}
}
