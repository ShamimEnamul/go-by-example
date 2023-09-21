package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

type dto struct {
	name    string
	age     int
	address string
}
type model struct {
	name    string
	age     int
	address string
}

func main() {
	a := dto{
		name:    "shamim",
		age:     14,
		address: "housebuilding",
	}

	var b model
	//b = model(a)
	copier.Copy(&a, &b)
	//err := copyutil.Copy(&a, &b)
	//if err != nil {
	//	fmt.Printf(err)
	//	return
	//}
	fmt.Println("b", b)

}
