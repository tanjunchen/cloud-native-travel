package main

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

var db struct {
	Dns string
}

func TestMain(m *testing.M) {

	db.Dns = "BBBB"
	if db.Dns == "" {
		db.Dns = "AAA"
	}

	flag.Parse()
	exitCode := m.Run()

	db.Dns = ""

	// 退出
	os.Exit(exitCode)
}

func TestDatabase(t *testing.T) {
	fmt.Println(db.Dns)
}
