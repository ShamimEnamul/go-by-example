package main

import "fmt"

type Day int

// Define constants for the enum values
const (
	Monday Day = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func String(d Day) {
	fmt.Println(d)
}

func main() {
	// String(Friday)
	var i interface{} = 1

	// Type assertion to get the underlying string
	//s := i.(string)
	//fmt.Println(s) // Output: hello

	// Safe type assertion using comma-ok idiom
	s, ok := i.(string)
	if ok {
		fmt.Println("String value:", s) // Output: String value: hello
	} else {
		fmt.Println("Not a string")
	}

	// Unsafe type assertion (will panic if incorrect)
	//n := i.(int) // This will panic because i does not hold an int
	//fmt.Println(n)

	val, ok := i.(int)
	fmt.Println(val)
}
