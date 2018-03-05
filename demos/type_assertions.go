package main

import (
	"log"
)

func main() {
	a := ttt()        // a为interface类型
	c, ok := a.(*aaa) // 此处必须加类型检查，如ok为true，c变量将是aaa类型
	if ok != false {
		log.Printf("%v", c.str)
	}
}

type aaa struct {
	str string
}

var fff = &aaa{str: "xxx"}

func ttt() interface{} {
	return fff
}
