package main

import (
	"bufio"
	"net"
	"os"
	"strings"
	"testing"
	"time"
)

func TestTcpEchoServer(t *testing.T) {
	prefix := "hello"
	request := "world"
	want := prefix + " " + request

	os.Args = []string{"main", "9000", prefix}
	go main()

	time.Sleep(2 * time.Second)

	for _, addr := range []string{":9000"} {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			t.Fatalf("couldn't connect to the server: %v", err)
		}
		defer conn.Close()

		if _, err := conn.Write([]byte(request + "\n")); err != nil {
			t.Fatalf("couldn't send request: %v", err)
		} else {
			reader := bufio.NewReader(conn)
			if response, err := reader.ReadBytes(byte('\n')); err != nil {
				t.Fatalf("couldn't read server response: %v", err)
			} else if !strings.HasPrefix(string(response), want) {
				t.Errorf("output doesn't match, wanted: %s, got: %s", want, response)
			}
		}
	}
}
