package main

import (
	"fmt"
	"time"
)

var l bool = false
var amount int = 100

func f1(x int) {
	fmt.Println(" f1 start")
	for l {
		fmt.Print(".")
	}
	l = true
	time.Sleep(3 * time.Second)
	amount = amount + 5
	l = false
	fmt.Println(" f1 end ")
}
func f2(x int) {
	fmt.Println(" f2 start")
	for l {
		fmt.Print(".")
	}
	l = true
	amount = amount - 5
	l = false
	fmt.Println(" f2 end ")
}

func main() {
	go f1(1)

	f2(1)

	time.Sleep(5 * time.Second)
	fmt.Println("Lock Mechanism")
	fmt.Println(l, "amount: ", amount)
}
