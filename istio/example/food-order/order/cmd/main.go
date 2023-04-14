package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Order")
}

func main() {
	fmt.Println("My name is Order")
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8888", nil)
}
