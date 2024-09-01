package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var sl []int
	//sl = append(sl, 2)
	//fmt.Println(sl)
	//
	//var sl = make([]int, 3)
	//sl[0] = 1
	//sl = append(sl, 2)

	// ### append ->  add element at end of the array. here len &  capacity = 3. After appending, [0,0,0,1]
	sl := make([]int, 3)
	sl = append(sl, 1)
	sl = append(sl, 2)
	sl = append(sl, 3)
	sl = append(sl, 4)
	//fmt.Println(sl)

	sl2 := make([]int, 3)
	sl[0] = 11
	// ### copy sl to sl2
	copy(sl2, sl)
	// ### assign sl into sl2. both will sl
	//sl2 = sl

	//fmt.Println("sl: ", sl)
	//fmt.Println("sl2: ", sl2)
	//fmt.Printf("sl: %p", sl)
	//fmt.Printf(" sl2: %p", sl2)

	for i, val := range sl2 {
		fmt.Println(i, ": ", val)
	}

	fmt.Println(reflect.DeepEqual(sl, sl2))
}
