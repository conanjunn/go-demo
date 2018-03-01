package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	userFile := "arr.go"
	fl, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}

	// 读取文件到byte slice中
	data, err := ioutil.ReadFile(userFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Data read: %v, %T\n", string(data), data)
}
