package main

import (
	"fmt"
	"time"
)

type Project struct{}

func (p *Project) deferError() {
	if err := recover(); err != nil {
		fmt.Println("recover: ", err)
	}
}

func (p *Project) exec(msgchan chan interface{}) {
	// 需要放到这里才能捕捉到 panic
	defer p.deferError()
	for msg := range msgchan {
		m := msg.(int)
		fmt.Println("msg: ", m)
	}
}

func (p *Project) run(msgchan chan interface{}) {
	for {
		// defer p.deferError() 需要在协程开始处调用，否则无法捕获 panic
		// defer p.deferError()
		go p.exec(msgchan)
		time.Sleep(time.Second * 2)
	}
}
func (p *Project) Main() {
	a := make(chan interface{}, 100)
	go p.run(a)
	go func() {
		for {
			// panic: interface conversion: interface {} is string, not int
			a <- "1"
			// a <- 1
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 10000000)
	// time.Second * 100000000000000 (constant 100000000000000000000000 of type time.Duration) overflows int64
	// time.Sleep(time.Second * 100000000000000)
}

func main() {
	p := new(Project)
	p.Main()
}
