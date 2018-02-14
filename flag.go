package main

import (
	"flag"
	"fmt"
)

func main() {
	// 拿命令行里的参数
	var port *int = flag.Int("port", 1000, "abc")
	flag.Parse()
	fmt.Println(*port, flag.Args())
}
