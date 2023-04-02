package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	//ctx := context.Background()
	//
	//ctxWithCancel, cancelFunction := context.WithCancel(ctx)
	//
	//defer func() {
	//	fmt.Println("Main Defer: canceling context")
	//	cancelFunction()
	//}()
	//
	//go func() {
	//	sleepRandom("Main", nil)
	//	cancelFunction()
	//	fmt.Println("Main Sleep complete. canceling context")
	//}()
	//
	//doWorkContext(ctxWithCancel)
	//test2()
	//test3()
	//test4()
	test5()
}

type Result struct {
	r   *http.Response
	err error
}

func test3() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan Result, 1)
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		fmt.Println("http request failed, err:", err)
		return
	}
	go func() {
		resp, err := client.Do(req)
		pack := Result{r: resp, err: err}
		c <- pack
	}()

	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		res := <-c
		fmt.Println("Timeout! err:", res.err)
	case res := <-c:
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("Server Response: %s", out)
	}
	return
}

func sleepRandom(fromFunction string, ch chan int) {
	defer func() { fmt.Println(fromFunction, "sleepRandom complete") }()
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	randomNumber := r.Intn(100)
	sleepTime := randomNumber + 100
	fmt.Println(fromFunction, "Starting sleep for", sleepTime, "ms")
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	fmt.Println(fromFunction, "Waking up, slept for ", sleepTime, "ms")
	if ch != nil {
		ch <- sleepTime
	}
}

func doWorkContext(ctx context.Context) {
	ctxWithTimeout, cancelFunction := context.WithTimeout(ctx, time.Duration(150)*time.Millisecond)

	defer func() {
		fmt.Println("doWorkContext complete")
		cancelFunction()
	}()

	ch := make(chan bool)
	go sleepRandomContext(ctxWithTimeout, ch)

	select {
	case <-ctx.Done():
		fmt.Println("doWorkContext: Time to return")
	case <-ch:
		fmt.Println("sleepRandomContext returned")
	}
}

func sleepRandomContext(ctx context.Context, ch chan bool) {
	defer func() {
		fmt.Println("sleepRandomContext complete")
		ch <- true
	}()
	sleepTimeChan := make(chan int)
	go sleepRandom("sleepRandomContext", sleepTimeChan)

	select {
	case <-ctx.Done():
		fmt.Println("sleepRandomContext: Time to return")
	case sleepTime := <-sleepTimeChan:
		fmt.Println("Slept for ", sleepTime, "ms")
	}
}

func test2() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	ctx = context.WithValue(ctx, "Test", "123456")
	defer cancelFunc()

	if t, ok := ctx.Deadline(); ok {
		fmt.Println(time.Now())
		fmt.Println(t.String())
	}
	go func(ctx context.Context) {
		fmt.Println(ctx.Value("Test"))
		for {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				return
			default:
				continue
			}
		}
	}(ctx)
	time.Sleep(time.Second * 3)
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("i exited")
				return // returning not to leak the goroutine
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func test() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers
	intChan := gen(ctx)
	for n := range intChan {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func test4() {
	test()
	time.Sleep(time.Hour)
}

func test5() {
	d := time.Now().Add(50 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
