package main

import "fmt"

func main() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)

	xxx := make([]int, 2, 32)

	fmt.Println(xxx)

	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl4 := append(sl3, 4, 5, 6)
	fmt.Println(sl3, sl4)
}
