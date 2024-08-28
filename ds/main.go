package main

import "fmt"

func main() {
	var arr [10]int

	arr[0] = 1
	//arr[100000000] = 10
	fmt.Println(len(arr), arr)

}
