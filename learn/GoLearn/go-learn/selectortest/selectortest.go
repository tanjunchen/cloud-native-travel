package selectortest

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) string {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)
	if aDuration < bDuration {
		return a
	}
	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func Racer2(urlA, urlB string) (winner string, err error) {
	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", urlA, urlB)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
