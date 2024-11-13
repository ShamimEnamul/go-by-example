package main

import "fmt"

type p struct {
	name string
}

type s struct {
	p
	age int
}

func main() {

	a := s{
		p: p{
			name: "shamim",
		},
		age: 10,
	}

	fmt.Println(a)
}
