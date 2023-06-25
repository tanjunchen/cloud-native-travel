package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	dataChan := make(chan int)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			dataChan <- rand.Intn(5)
		}
		close(dataChan)
	}(&waitGroup)
	go func(wg *sync.WaitGroup) {
		defer waitGroup.Done()
		for i := range dataChan {
			fmt.Println(i)
		}
	}(&waitGroup)
	waitGroup.Wait()
}
