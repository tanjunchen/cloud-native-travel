package Chapter01

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

func main018() {
	for _, link := range os.Args[4:] {
		if !strings.HasPrefix(link, "https://") && !strings.HasPrefix(link, "http://") {
			link = "https://" + link
		}

		resp, err := http.Get(link)
		if err != nil {
			log.Fatal(err)
		}
		bytes, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("%s", bytes)
	}
}

// go test -v -run  Test018 practice1.8_test.go -args api.jt-gmall.com
func Test018(t *testing.T) {
	main018()
}
