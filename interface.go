package main

import "fmt"

type Person interface {
	walk(step int) int
}

func main() {
	var laoshi Person
	laoshi = &Teacher{name: "laoshi"}
	// laoshi := &Teacher{name: "laoshi"} // 等价于上面的写法

	xuesheng := &Student{name: "xxx"}

	doWalk(laoshi)
	doWalk(xuesheng)
}

func doWalk(people Person) {
	people.walk(1)
	people.walk(2)
	people.walk(3)
	people.walk(4)
}

type Teacher struct {
	name string
}

func (this *Teacher) walk(step int) int {
	fmt.Printf("%s walk %d step\n", this.name, step)
	return step
}

type Student struct {
	name string
}

func (this *Student) walk(step int) int {
	fmt.Printf("%s walk %d step\n", this.name, step)
	return step
}
