package main

import "fmt"

func main() {
	a, b := abc(1, "aaa")
	fmt.Println(a, b)

	var arr []string = []string{"a", "b"}
	c, d := demo(1, arr...)
	fmt.Println(c, d)

	callback(1, func(a, b int) int {
		fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
		return 1
	})
}

func abc(a int, b string) (int, string) {
	return a, b
}

// defer & 变长参数
func demo(a int, b ...string) (int, string) {
	var str string = ""
	defer func() {
		fmt.Println("defer")
	}()
	for i := 0; i < len(b); i++ {
		str += (b[i] + "|")
	}
	return a, str
}

func callback(y int, f func(int, int) int) { // f 参数的返回类型和参数类型必须和传入说的函数的一模一样
	a := f(y, 2)
	fmt.Println(a)
}
