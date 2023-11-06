package main

import (
	"bytes"
	"net/http"
	"time"
)

var payload = bytes.Repeat([]byte("0"), 1024)

func handler(w http.ResponseWriter, req *http.Request) {
	sleep := false
	if req.Header.Get("sleep") == "true" {
		sleep = true
	}
	if sleep {
		time.Sleep(5 * time.Millisecond)
		w.Write(payload)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
