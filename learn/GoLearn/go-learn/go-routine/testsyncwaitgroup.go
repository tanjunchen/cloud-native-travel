package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func run(task string) {
	fmt.Println(task, "start ....")
	time.Sleep(time.Second * 2)
	wg.Done()
}

func test11() {
	wg.Add(2)

	for i := 1; i < 3; i++ {
		taskName := "task " + strconv.Itoa(i)
		go run(taskName)
	}

	wg.Wait()
	fmt.Println("over")
}

func test12() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Task 1 ... over ")
			default:
				fmt.Println("Task 1 running ... ")
				time.Sleep(time.Second * 2)
			}
		}
	}()

	time.Sleep(time.Second * 10)
	fmt.Println("stop Task 1")
	stop <- true
	time.Sleep(time.Second * 3)
}

func main() {
	test12()
}
