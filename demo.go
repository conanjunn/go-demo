package main

import "fmt"

func main() {
	var arr = []int{0, 1, 2, 3, 4}
	var arr1 = sum(arr[:])
	fmt.Println(arr1)
}

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}
