package main

import "fmt"

func handlePainc() {
	fmt.Println("handle function called!")
	r := recover()
	if r != nil {
		fmt.Printf("panic occured")
	}
}
func main() {
	defer handlePainc()
	panic("panic raised")
	fmt.Println("Hello world")
}
