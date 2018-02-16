package main

import "fmt"

func main() {
	var i int = 1
	var ptr *int = &i // 取址

	fmt.Printf("值为：%d，址为：%p\n", i, ptr)

	var ii int = *ptr // 取值

	fmt.Printf("value: %d, point: %p\n", ii, &ii)

	*ptr = 2
	fmt.Printf("old value: %d, new value: %d\n", i, ii)

	// 常量不能取址
	// const a = 1
	// &a 报错
	// &1 报错
}
