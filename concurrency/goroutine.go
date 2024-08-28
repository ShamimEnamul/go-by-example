package main

import (
	"fmt"
	"time"
)

func start() {
	go start2()
	fmt.Println("In Goroutine!")
}

func start2() {
	fmt.Println("In Goroutine2!")
}

func execute(id int) {
	fmt.Printf("id: %d\n", id)
}

func example1() {
	go start()
	fmt.Println("started!")
	time.Sleep(1 * time.Second)
	fmt.Println("Finished!")
}

// example with loop
func example2() {
	fmt.Println("started!")
	for i := 0; i < 10; i++ {
		go execute(i)
	}
	time.Sleep(2 * time.Second) // if we remove the timeout then main will exit, so nothing will be executed from execute file
	fmt.Println("Finished!")
}
