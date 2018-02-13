package main

import (
	"flag"
	"fmt"
)

func main() {
	var port *int = flag.Int("port", 1000, "abc")
	flag.Parse()
	fmt.Println(*port, flag.Args())
}
