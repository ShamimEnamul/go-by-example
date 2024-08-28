package main

import (
	"fmt"
)

type Orange struct {
	Quantity int
}

func (o *Orange) Increase(n int) {
	o.Quantity += n
}

func (o *Orange) Decrease(n int) {
	o.Quantity -= n
}

func (o *Orange) String() string {
	fmt.Println("String called!")
	return fmt.Sprintf("%v", o.Quantity)
}

func main() {
	var orange Orange

	orange.Increase(10)
	orange.Decrease(5)
	// p := &orange
	fmt.Println(orange)
	//p.String()
	//orange.String()
	//cc := orange.String()
	//io.WriteString(os.Stdout, cc)

	// ...........................
	// Compare interface
	//var a interface{}
	//var b interface{}
	//
	//a = 1
	//b = 2
	//fmt.Println(a == b)
	//
	//a = []int{1, 2, 3}
	//b = []int{1, 2, 3}
	//// fmt.Println(a == b) // not possible
	//fmt.Println(reflect.DeepEqual(a, b))
	// .............................
	// Compare struct
	//type Foo struct {
	//	A int
	//	B int
	//	C interface{}
	//}
	//
	//a := Foo{
	//	A: 1,
	//	B: 2,
	//	C: 3,
	//}
	//
	//b := Foo{
	//	A: 1,
	//	B: 2,
	//	C: 3,
	//}

	//type Foo struct {
	//	A int
	//	B int
	//	C []int
	//}
	//
	//a := Foo{
	//	A: 1,
	//	B: 2,
	//	C: []int{1, 2, 3},
	//}
	//
	//b := Foo{
	//	A: 1,
	//	B: 2,
	//	C: []int{1, 2, 3},
	//}
	//
	//fmt.Println(reflect.DeepEqual(a, b))
	// or
	// loop through the array defined in struc

	// .......................

	// Copy or assign map
	//a := map[string]bool{"A": true, "B": false}
	//b := map[string]bool{"C": true, "D": false}
	////a = b
	//fmt.Println("A:", a, " B:", b)
	//
	//for key, value := range a {
	//	b[key] = value
	//}
	//// copy all key-value from A
	//fmt.Println("A:", a, " B:", b)

	// copy or assign slice
	//a := []int{1, 2, 3}
	//b := []int{4, 5, 6}
	//
	//check := a
	////a = b
	//copy(a, b)
	//
	//fmt.Println(check)

	// swap
	//a, b := 1, 2
	//a, b = b, a
	//fmt.Printf("%d %d", a, b)
}
