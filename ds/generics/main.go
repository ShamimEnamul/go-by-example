package main

import (
	"fmt"
)

type a = int

type Dog[T any] struct {
	name string
	val  T
	d    *Dog[T]
}

func sum[T any](val T) {
	fmt.Println(val)
}

func sub[S ~[]E, E comparable](s S, val E) {
	fmt.Println(s, val)
}
func main() {
	sum("afaf")
	sum(12)

	d := Dog[int]{
		name: "shamim",
		val:  12,
		d:    &Dog[int]{},
	}
	fmt.Println(d)

}
