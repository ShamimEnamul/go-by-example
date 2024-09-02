package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println("RECURSION")
	fmt.Println(fact(5))

	var f func(int) int

	f = func(n int) int {
		if n == 0 {
			return 1
		}
		return n * f(n-1)
	}

	fmt.Println("closure: ", f(5))
}
