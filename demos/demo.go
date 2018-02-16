package main

import (
	"fmt"
)

func demo(s []int) {
	s = append(s, 3)
}

func main() {
	s := make([]int, 1)
	fmt.Println(s)
	demo(s)
	fmt.Println(s)
}
