package main

import (
	"fmt"
)

func main() {
	fmt.Println("MAP")

	m := make(map[string]int)
	m["shamim"] = 12
	m["s"] = 102
	m["ss"] = 103
	m["sss"] = 104
	fmt.Println(m, "len", len(m))
	val := m["shamim"]

	fmt.Println("value", val)
	isAvailable, v := m["shamim"]

	fmt.Println(" isAvailable represents a key:'shamim' exist in the map or not", "v represents the Value of the key 'shamim'")
	fmt.Println(isAvailable, " : ", v)

	delete(m, "shamim")
	fmt.Println("len: ", len(m))

	for key, value := range m {
		fmt.Println(key, " : ", value)
		//fmt.Println(m[key])
	}

}
