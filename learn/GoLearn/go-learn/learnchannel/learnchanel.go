package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(name string, channel chan string) {
	for {
		channel <- fmt.Sprintf("%s: %v", name, rand.Int31())
		time.Sleep(time.Second)
	}
}

func consumer(channel chan string) {
	for {
		message := <-channel
		fmt.Println(message)
	}
}

func test2() {
	name := "张三"
	fmt.Println(name)
	modify(&name)
	fmt.Println(name)
}

func modify(s *string) {
	*s = *s + *s
}

func test3() {
	var arrayA = [3]string{"hammer", "soldier", "mum"}

	for _, value := range arrayA {
		fmt.Println(value)
	}
}

type severity int32 // sync/atomic int32

const (
	infoLog severity = iota
	warningLog
	errorLog
	fatalLog
	numSeverity = 4
)

func main() {
	//var channel chan string = make(chan string)
	//go producer("xxxx", channel)
	//go producer("****", channel)
	//
	//consumer(channel)
	//test3()
	fmt.Println(infoLog, warningLog, errorLog, fatalLog, numSeverity)
}
