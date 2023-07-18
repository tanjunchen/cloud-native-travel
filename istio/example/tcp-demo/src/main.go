package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

var (
	port   string
	suffix string
)

func init() {
	flag.StringVar(&port, "port", "8000", "Port for Server")
	flag.StringVar(&suffix, "suffix", "TCP", "The Suffix of Response")
}

func main() {
	args := os.Args
	var ports []string
	var s string

	if len(args) > 1 {
		ports = strings.Split(os.Args[1], ",")
	} else {
		ports = append(ports, port)
	}

	if len(args) > 2 {
		s = os.Args[2]
	} else {
		s = suffix
	}

	for _, port := range ports {
		addr := fmt.Sprintf(":%s", port)
		go serve(addr, s)
	}
	ch := make(chan struct{})
	<-ch
}

func serve(addr, prefix string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("failed to create listener, err:", err)
		os.Exit(1)
	}
	fmt.Printf("listening on %s, prefix: %s\n", listener.Addr(), prefix)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept connection, err:", err)
			continue
		}

		go handleConnection(conn, prefix)
	}
}

func handleConnection(conn net.Conn, prefix string) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				fmt.Println("failed to read data, err:", err)
			}
			return
		}
		fmt.Printf("request: %s", bytes)

		line := fmt.Sprintf("%s %s", prefix, bytes)
		fmt.Printf("response: %s", line)
		conn.Write([]byte(line))
	}
}
