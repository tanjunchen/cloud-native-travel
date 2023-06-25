package main

import (
	"fmt"
	"runtime"
)

// 这个跟 golang 的版本有关，高版本 golang 会打印出 Dropping mic 与 Done
func main() {
	var i byte
	go func() {
		// i<=255  256 (untyped int constant) overflows byte
		for i = 0; i <= 255; i++ {
			fmt.Println(i)
		}
	}()
	fmt.Println("Dropping mic")
	// Yield execution to force executing other goroutines
	runtime.Gosched()
	runtime.GC()
	fmt.Println("Done")
}
