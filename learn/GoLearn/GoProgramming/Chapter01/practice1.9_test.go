package Chapter01

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

func main019() {
	for _, link := range os.Args[4:] {
		if !strings.HasPrefix(link, "https://") && !strings.HasPrefix(link, "http://") {
			link = "https://" + link
		}

		resp, err := http.Get(link)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		fmt.Printf("%v", resp.StatusCode)
	}
}

// go test -v -run  Test019 practice1.9_test.go -args api.jt-gmall.com
func Test019(t *testing.T) {
	main019()
}
