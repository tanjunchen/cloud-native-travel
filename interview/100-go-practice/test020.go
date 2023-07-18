package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
*
2---i:  9
1---i:  10
1---i:  10
1---i:  10
1---i:  10
1---i:  10
1---i:  10
1---i:  10
1---i:  10
1---i:  10
1---i:  10
2---i:  0
2---i:  1
2---i:  2
2---i:  3
2---i:  4
2---i:  5
2---i:  6
2---i:  7
2---i:  8
*
*/
func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("1---i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("2---i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
