package main

import "fmt"

// variadic function
func sum(nums ...int) {
	for i := range nums {
		fmt.Println(nums[i])
	}
}
func main() {
	sum(1, 2, 3)
	arr := []int{5, 6, 7}

	sum(arr...)
}
