package main

import "fmt"

// 切片本身已经是一个引用类型，所以它本身就是一个指针!!

func main() {
	var arr1 [6]int

	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	fmt.Printf("arr1 is %d\n", arr1)

	var slice1 []int = arr1[0:6] // item at index 5 not included!

	fmt.Printf("slice1 is %d\n", slice1)
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
}
