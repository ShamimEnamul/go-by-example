package main

import (
	"fmt"
	"golang/package/calculator"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := calculator.Average(xs)
	fmt.Println(avg)
}
