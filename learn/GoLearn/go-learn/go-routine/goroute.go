package main

import "fmt"

//function to print hello
func printHello() {
	fmt.Println("Hello from printHello")
}

func test1() {
	//inline goroutine. Define a function inline and then call it.
	go func() { fmt.Println("Hello inline") }()
	//call a function as goroutine
	go printHello()
	fmt.Println("Hello from main")
}

func printHello2(ch chan int) {
	fmt.Println("Hello from printHello")
	//send a value on channel
	ch <- 2
}
func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("Hello inline")
		//send a value on channel
		ch <- 1
	}()

	go printHello2(ch)
	fmt.Println("Hello from main")

	i := <-ch
	fmt.Println("Recieved ", i)

	<-ch
}
