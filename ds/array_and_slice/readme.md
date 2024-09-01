# Array

### array size declaration
	var arr [5]int
	arr[2] = 22
	fmt.Println("arr: ", arr)

### array with initialization
	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr2: ", arr2)

    
### Let compiler count the length of array
	var arr4 = [...]int{1, 2, 3}
	fmt.Println(len(arr4))
    
    arr4 = [...]

### insert value with index using :
	var arr4 = [...]int{1, 2, 3}
	fmt.Println("arr4: ", arr4)
	arr4 = [...]int{1: 2, 2: 3}

	fmt.Println("arr4 with index input: ", arr4)



# Slice


## slice with initialization and also can append
	var arr3 = []int{1, 2, 3}
	fmt.Println("arr3", arr3)
	arr3 = append(arr3, 21)

