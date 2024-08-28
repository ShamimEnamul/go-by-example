package main

import "fmt"

type Stack []int
type a struct {
	ss int
}

func (s *Stack) Push(int val) {
	s = append(*s, val)
	fmt.Println(s)

	//p := a{
	//	ss: 10,
	//}
	//l := &p
	//(*l).ss = 100

}

func Pop() {

}

func Peek() {

}

func Empty() bool {
	return false
}
