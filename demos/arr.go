package main

import "fmt"

func main() {
	a := [...]string{"a", "b", "c", "d"} // 逗号可以省略
	b := [5]int{1, 2, 3, 4, 5}
	c := [5]string{3: "m", 4: "n"}

	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

	var arr1 = new([5]int) // arr1 为指针

	fmt.Println("a:", a)
	fmt.Println("arr", arr1, b, c)
}
