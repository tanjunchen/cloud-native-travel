package main

import (
	"fmt"
	"runtime"
)

func main() {
	/**
	panic: hello

	goroutine 1 [running]:
	main.main()
	        /Users/chentanjun/opensource/cloud-native-travel/interview/100-go-practice/test022.go:18 +0x120
	exit status 2
	*/
	/**
	结果是随机执行。golang 在多个 case 可读的时候会公平的选中一个执行。
	*/
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}
