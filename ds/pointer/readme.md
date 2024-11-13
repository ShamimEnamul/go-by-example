	

    // storing the hexadecimal
	// values in variables
	x := 0xFF
	y := 0x9C

	// Displaying the values
	fmt.Printf("Type of variable x is %T\n", x)
	fmt.Printf("Value of x in hexadecimal is %X\n", x)
	fmt.Printf("Value of x in decimal is %v\n", x)

	fmt.Printf("Type of variable y is %T\n", y)
	fmt.Printf("Value of y in hexadecimal is %X\n", y)
	fmt.Printf("Value of y in decimal is %v\n", y)


### declare and assign address to a pointer
	var x int = 100

	var p *int

	p = &x

	fmt.Println("Value of x ", x)
	fmt.Println("Address of x ", &x)
	fmt.Println("Value stored in p ", p)

### The default value of a pointer in <nil>
var x *int
fmt.Println(x)