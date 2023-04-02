package Chapter01

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

func TimeConsuming2(tag string) func() {
	now := time.Now().UnixNano() / 1000000
	return func() {
		after := time.Now().UnixNano() / 1000000
		fmt.Printf("%q time cost %d ms\n", tag, after-now)
	}
}

func main0110() {
	ch := make(chan string)
	for _, link := range os.Args[4:] {
		go fetch(link, ch)
	}
	for range os.Args[4:] {
		fmt.Println(<-ch)
	}
}

func fetch(link string, ch chan<- string) {
	defer TimeConsuming2("Fetch" + link)
	if !strings.HasPrefix(link, "https://") && !strings.HasPrefix(link, "http://") {
		link = "https://" + link
	}

	resp, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	ch <- string(bytes)
}

// go test -v -run  Test0110 practice1.10_test.go -args api.jt-gmall.com
func Test0110(t *testing.T) {
	main0110()
}
