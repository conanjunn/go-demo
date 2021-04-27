package main

import "fmt"

// 模拟class
type Animal struct {
	name string
	age  int
}

// 接受器为指针的方法
func (this *Animal) run() {
	this.age++ // 会使原值得age加1
}

// 接收器为值得方法
func (this Animal) run1() {
	this.age++ // 仅会使此函数中的this上的age加1，this是原值得一个copy
}

func main() {
	an := &Animal{ // an为指针
		name: "abc",
		age:  1,
	}
	/*
		等同于上面的写法
		an := new(Animal)
		an.age = 1
		an.name = "aaa" */

	an.run()
	fmt.Printf("这里age为2： %+v \n", an)

	an.run1()
	fmt.Printf("这里age依然为2： %+v \n", an)
}
