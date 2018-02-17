package main

import (
	"errors"
	"fmt"
)

var errNotFound error = errors.New("Not found error \n")

func test() (num int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()
	panic("no words to parse")
	num = 1
	return
}

func main() {
	fmt.Printf("error: %v \n", errNotFound)
	num, err := test()
	fmt.Printf("test complete, %d, %v \n", num, err)
}
