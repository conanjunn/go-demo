package main

import "fmt"

func main() {
	var bl bool // 必须是bool类型
	if bl {
		fmt.Println("if")
	} else if i := 1; i == 1 {
		fmt.Println("else if")
	} else {
		fmt.Println("else")
	}
}
