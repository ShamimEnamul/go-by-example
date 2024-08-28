package main

import (
	"fmt"
	"time"
)

func counter1(str chan string) {
	time.Sleep(5 * time.Second)
	str <- "counter 1"
	return
}

func counter2(str chan string) {
	time.Sleep(3 * time.Second)
	str <- "counter 2"
	return
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go counter1(c1)
	go counter2(c2)
	for i := 0; i < 2; i++ {

		select {
		case val := <-c1:
			fmt.Println(val)
		case val := <-c2:
			fmt.Println(val)
		default:
			fmt.Println("default called!")
		}
	}
}
