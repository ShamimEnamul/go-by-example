package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	log.Println("booting on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	//fmt.Println("RECURSION")
	//fmt.Println(fact(5))
	//
	//var f func(int) int
	//
	//f = func(n int) int {
	//	if n == 0 {
	//		return 1
	//	}
	//	return n * f(n-1)
	//}
	//
	//fmt.Println("closure: ", f(5))
}
