package main

import (
	"bufio"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	prefix := "TCP"
	request := "TCP TEST example"
	want := prefix + " " + request

	time.Sleep(2 * time.Second)

	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		log.Printf("couldn't connect to the server: %v", err)
	}
	defer conn.Close()

	if _, err := conn.Write([]byte(request + "\n")); err != nil {
		log.Printf("couldn't send request: %v", err)
	} else {
		reader := bufio.NewReader(conn)
		if response, err := reader.ReadBytes(byte('\n')); err != nil {
			log.Printf("couldn't read server response: %v", err)
		} else if !strings.HasPrefix(string(response), want) {
			log.Printf("output doesn't match, wanted: %s, got: %s", want, response)
		}
	}
}
