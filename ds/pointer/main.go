package main

import "fmt"

func main() {

	var x = 10
	var y = &x
	fmt.Println(y)
	fmt.Println("value of the x pointing by y: ", *y)
}
