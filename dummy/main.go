package main

import (
	"fmt"
)

type Animal struct {
	name string
}

func (a *Animal) eat() {
	fmt.Println(a.name + " eating!!!")
}

func (a *Animal) walk() {
	fmt.Println("Walking!!!")
}

type Dog struct {
	Animal
}

func main() {
	//c1 := make(chan string, 1)
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	c1 <- "result 1"
	//}()
	//
	//select {
	//case res := <-c1:
	//	fmt.Println(res)
	//case <-time.After(1 * time.Second):
	//	fmt.Println("timeout 1")
	//}
	//
	//c2 := make(chan string, 1)
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	c2 <- "result 2"
	//}()
	//select {
	//case res := <-c2:
	//	fmt.Println(res)
	//case <-time.After(3 * time.Second):
	//	fmt.Println("timeout 2")
	//}

	//arr := [...]int{1, 2, 3}
	//arr[0] = 1
	//fmt.Println(arr)
	//var ar []int
	//ar = append(ar, 2)
	//// ar[0] = 20
	//fmt.Println(len(ar))

	mp := make(map[int]int)

	mp[0] = 1
	mp[2] = 2
	fmt.Println(mp)
	val, isExit := mp[0]
	delete(mp, 0)
	fmt.Println(mp)
	fmt.Println(val, isExit)
	a, b := 1, 2
	fmt.Println(a, b)

}
