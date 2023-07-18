package main

import (
	"fmt"
	"sync/atomic"
)

var value int32

func SetValue(delta int32) {
	for {
		v := value
		fmt.Println("Before CompareAndSwapInt32", v, delta)
		if atomic.CompareAndSwapInt32(&value, v, (v + delta)) {
			fmt.Println("CompareAndSwapInt32", v, delta)
			break
		}
		fmt.Println("After CompareAndSwapInt32", v, delta)
	}
}

// atomic.CompareAndSwapInt32 函数不需要循环调用。
func main() {
	value = 10
	SetValue(2)
}
