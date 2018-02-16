package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	chanFn()
	rangeChan()
	fmt.Println("---------------------", runtime.NumCPU())
	multiChan()
	multiChan2()
	bufferChan()
	demo()
}

func chanFn() {
	// chan基本用法
	ch := make(chan int)
	go func() {
		ch <- 11
	}()
	var i int
	i = <-ch
	fmt.Printf("result: %d\n", i)
}

func rangeChan() {
	// 迭代chan
	ch1 := make(chan int)
	go func() {
		ch1 <- 123
		ch1 <- 456
		ch1 <- 789
		close(ch1) // 这里必须主动关闭
	}()
	for v := range ch1 {
		fmt.Printf("for range: %d\n", v)
	}
}

// 带缓冲的chan解决多个goroutine并发
func multiChan() {
	ch := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go Go(ch, i)
	}
	fmt.Println("is gogo")

	for i := 0; i < 10; i++ {
		bl := <-ch
		fmt.Printf("bl: %v \n", bl)
	}
}

func Go(ch chan bool, i int) {
	fmt.Println("go go", i)
	ch <- true
}

// waitGroup解决多个goroutine并发
func multiChan2() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Go2(&wg, i)
	}
	wg.Wait()
}

func Go2(wg *sync.WaitGroup, i int) {
	fmt.Println("wg", i)
	wg.Done()
}

// 带缓存的chan
func bufferChan() {
	ch := make(chan int, 3)
	ch <- 1
	fmt.Println(<-ch)
	ch <- 2
	ch <- 3
	ch <- 4
	close(ch) // 必须显示关闭
	for v := range ch {
		fmt.Println(v)
	}
}

func demo() {
	fmt.Println("demo-----------------")
	ch := make(chan int, 3)
	go func() {
		fmt.Println("gogog")
		ch <- 1
	}()
	fmt.Println("ddddddd")
	<-ch
}
