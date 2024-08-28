package main

import (
	"fmt"
	"time"
)

// buffered channel

func send(str chan string) {
	str <- "hello"
	fmt.Println("sending")

}

func receive(str chan string) {
	time.Sleep(1 * time.Second)
	fmt.Println("received: ", <-str)
}

func sum(ch chan int) {
	sum := 0
	for val := range ch {
		sum += val
	}
	fmt.Printf("Sum: %d\n", sum)
}
func main() {
	ch := make(chan int, 3)
	ch <- 2
	ch <- 2
	ch <- 2
	close(ch)
	go sum(ch)

	//c1 := make(chan string, 1)
	//
	//// close channel
	//c1 <- "val"
	//close(c1)
	//time.Sleep(1 * time.Second)
	//val, ok := <-c1
	//fmt.Println("Received1: ", val, ok)

	//capacity of a channel
	//fmt.Println(cap(c1))
	// waiting for the message
	//c1 <- "val1"
	//fmt.Println("sending")
	//fmt.Println("Received1: ", <-c1)
	//fmt.Println("Received2: ", <-c1)

	// populate the channeld more than it's capacity
	//c1 <- "val1"
	//c1 <- "val2"
	//fmt.Println("sending")
	//fmt.Println("Received: ", <-c1)

	//
	//go send(c1)
	//go send(c1)
	//go send(c1)
	//go send(c1)
	//go send(c1)
	//go receive(c1)
	//go receive(c1)
	//go receive(c1)
	time.Sleep(4 * time.Second)
}

// unbuffered channel
//func send(str chan string) {
//	time.Sleep(1 * time.Second)
//	fmt.Println("Sending the message")
//	str <- "hello"
//	// time.Sleep(1 * time.Second)
//}
//
//func receive(str chan string) {
//	fmt.Println("Received: ", <-str)
//	fmt.Println("Receiving the message")
//}
//
//func main() {
//	c1 := make(chan string)
//	go receive(c1)
//	go send(c1)
//	time.Sleep(2 * time.Second)
//}
