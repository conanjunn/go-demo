package main

import (
	"fmt"
	"reflect"
)

type Person interface {
	info()
}

type User struct {
	str string
	i   int
}

func (this *User) info() {
	fmt.Printf("aaa")
}

func main() {
	var user interface{}  // 必须是空接口
	user = User{"aaa", 1} // 不可以用&  必须把struct属性全部初始化
	ty := reflect.TypeOf(user)
	val := reflect.ValueOf(user)
	fmt.Println("type: ", ty, "value: ", val, "kind:", ty.Kind())

	fmt.Println("NumField", ty.NumField(), val.NumField())

	for i := 0; i < ty.NumField(); i++ {
		f := ty.Field(i)
		fmt.Println("key:", f.Name, "type:", f.Type)

		fmt.Println("value fried: ", val.Field(i))
	}

	fmt.Println("NumMethod", ty.NumMethod())

	for i := 0; i < ty.NumMethod(); i++ {
		m := ty.Method(i)
		fmt.Println("method: ", m.Name)
	}
}
