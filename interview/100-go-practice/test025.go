package main

import (
	"fmt"
	"sync"
	"time"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}
func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	// map属于引用类型，并发读写时多个协程⻅是通过指针访问同一个地址，即访问共享变量，此时同时读写资源存在竞争关系。
	// 会报错误信息:“fatal error:concurrent map read and map write”。
	userAges := &UserAges{
		ages:  map[string]int{},
		Mutex: sync.Mutex{},
	}
	go func(userAges *UserAges) {
		for i := 0; i < 10; i++ {
			userAges.Add("tom", i)
		}
	}(userAges)
	go func(userAges *UserAges) {
		for i := 0; i < 10; i++ {
			fmt.Println(userAges.Get("tom"))
		}
	}(userAges)
	time.Sleep(time.Second * 5)
}
