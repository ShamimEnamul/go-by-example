package main

import (
	"fmt"
	"sync"
	"time"
)

func out() {
	for i := 0; i < 5; i++ {
		fmt.Println(" world! ", i)
		time.Sleep(time.Millisecond * 500)
	}
}
func print(thing string, c chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(thing, i)
		c <- "printed"
		time.Sleep(time.Millisecond * 500)
	}

	close(c)

}
func isEven(n int) bool {
	return n%2 == 0
}
func main() {
	//wg := sync.WaitGroup{}
	//fmt.Println(runtime.NumCPU())
	//wg.Add(1)
	//go func() {
	//	print("hello")
	//	wg.Done()
	//}()
	//
	//wg.Wait()
	//
	////print("world")

	//c1 := make(chan string)
	//c2 := make(chan string)
	//
	//go func() {
	//	for {
	//		c1 <- "c1"
	//		time.Sleep(time.Millisecond * 500)
	//	}
	//
	//}()
	//
	//go func() {
	//	for {
	//		c2 <- "c2"
	//		time.Sleep(time.Second * 2)
	//	}
	//
	//}()
	//for {
	//	select {
	//	case msg := <-c1:
	//		fmt.Println(msg)
	//	case msg := <-c2:
	//		fmt.Println(msg)
	//	}
	//	fmt.Println(<-c1)
	//	fmt.Println(<-c2)
	//}

	n := 0
	var m sync.Mutex
	//var m1 sync.Mutex

	// now, both goroutines call m.Lock() before accessing `n`
	// and call m.Unlock once they are done
	go func() {
		m.Lock()
		defer m.Unlock()
		nIsEven := isEven(n)
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, "is odd")
	}()

	go func() {
		m.Lock()
		n++
		m.Unlock()
	}()

	time.Sleep(time.Second)

}
