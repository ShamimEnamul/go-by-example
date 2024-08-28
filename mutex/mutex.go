package main

import (
	"fmt"
	"time"
)

func add(n int) int {
	fmt.Println("add called!")
	return n + 1
}
func print(n int) {
	//time.Sleep(2 * time.Second)
	fmt.Println("n ", n)
}
func main() {
	n := 1
	//var m sync.Mutex

	go func() {
		//m.Lock()
		//defer m.Unlock()
		time.Sleep(100 * time.Millisecond)
		fmt.Println(".....")
		print(n)
	}()
	fmt.Println("before add call()")
	go add(n)
	time.Sleep(2 * time.Second)

}
