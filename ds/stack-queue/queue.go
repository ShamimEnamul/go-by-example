package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// fmt.Println("enter number")
	//var val int
	//fmt.Scanln(&val)
	//fmt.Println("My age: ", val)
	//
	//var str string
	//fmt.Scanln(&str)
	//fmt.Println("My name is ", str)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter full address!")

	strr, _ := reader.ReadString('\n')
	fmt.Println(strr)

	scanner := bufio.NewScanner(strings.NewReader(`
	one
two
three
`))

	var (
		text []byte
		n    int
	)

	for scanner.Scan() {
		n++
		text = append(text, fmt.Sprintf("%d. %s\n", n, scanner.Text())...)
	}
	os.Stdout.Write(text)
	//var v int
	//fmt.Scanf("%d", &v)
	//fmt.Println("my age is ", v)

	//var name string
	//fmt.Print("Enter your name: ")
	//fmt.Scanln(&name)  // Take input until a newline is encountered
	//fmt.Println("Hello,", name)
	//stack := list.New()
	//
	//stack.PushBack(8)
	//stack.PushBack(9)
	//
	//fmt.Println(stack.Len())
	//stack.Remove(stack.Back())
	//e := stack.Front()
	//fmt.Println(e.Value)
}
