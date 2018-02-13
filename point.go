package main

import "fmt"

func main() {
	var i int = 1
	var ptr *int = &i // 取址

	fmt.Println(i, ptr)

	var ii int = *ptr // 取值

	fmt.Println(ii, &ii)

	*ptr = 2

	fmt.Println(i, ii)

	// const a = 1
	// &a 报错
	// &1 报错
}
