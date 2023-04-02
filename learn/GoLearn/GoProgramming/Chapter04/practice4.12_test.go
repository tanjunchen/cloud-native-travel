package Chapter03

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"
)

type Xvcd struct {
	Title      string
	Img        string
	Transcript string
}

func start0412() {
	var urls []string
	for i := 0; i < 1000; i++ {
		url := fmt.Sprintf("https://xkcd.com/%d/info.0.jsontest", i)
		urls = append(urls, url)
	}
	//var content []Xvcd
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch)
	}
	fmt.Println(<-ch)
}

func fetch(url string, ch chan<- string) {
	var result Xvcd
	fmt.Println(url)
	resp, _ := http.Get(url)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		os.Exit(1)
	}
	json.NewDecoder(resp.Body).Decode(&result)
	titles := strings.Split(result.Transcript, " ")

	for _, v := range titles {
		if v == os.Args[4] {
			ch <- result.Img
			break
		}
	}
}

// go test -v -run Test0412 practice4.12_test.go  -args woman
func Test0412(t *testing.T) {
	start0412()
}
