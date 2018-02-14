package main

import "fmt"

type primary struct {
	pri string
}

func main() {
	type struct1 struct {
		i1       int
		str      string
		*primary // 继承
	}

	var a = new(struct1) //  a为指针
	a.i1 = 1
	a.str = "aaa"
	var aa = &struct1{i1: 123, primary: &primary{pri: "kkk"}} // 等价于上面的new定义
	// aa.pri = "kkk"
	// aa.primary.pri = "kkk" // 等价于上面的写法
	fmt.Println("a是new出来的: ", a, aa, aa.GetPri())

	var b struct1 // b 为普通值 必须用var  :=不行
	b.i1 = 2
	fmt.Println("b是普通定义出来的: ", b)

	var c struct1 // c 为普通值
	fmt.Println("c是普通定义出来的: ", c)

	var pri *primary
	var result int = pri.Method()
	fmt.Println(result, "method called")
	// fmt.Println(pri.GetPri())  报错，传入的不是指针
}

func (v *primary) Method() int {
	return 100
}

func (v *primary) GetPri() string {
	return v.pri
}
