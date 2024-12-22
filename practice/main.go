package main

import "fmt"

type Eater interface {
	eat()
}
type Animal struct{}

func (a Animal) eat() {
	fmt.Println("Animal is eating...")
}

type Dog struct {
	Animal
}

func (d Dog) bark() {
	fmt.Println("Dog is barking...")
}

func main() {
	var eater Eater

	a := Animal{}
	eater = a
	eater.eat()

	fmt.Println("After using composition")
	d := Dog{}
	eater = d
	d.bark()
	eater.eat()

}
