package main

import "fmt"

func main() {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack)

	// pop
	stack.Pop()
	stack.Pop()
	fmt.Println(stack)

	// check empty
	fmt.Println(stack.Empty())

	// check peek
	fmt.Println(stack.Peek())
}
