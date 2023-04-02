package Chapter01

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
)

func main017() {
	for i, link := range os.Args[4:] {
		resp, err := http.Get(link)
		if err != nil {
			log.Fatal(err)
		}

		filename := fmt.Sprintf("file-test%d", i)
		out, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}

		if _, err := io.Copy(out, resp.Body); err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		out.Close()
	}
}

// go test -v -run  Test017 practice1.7_test.go -args "https://api.jt-gmall.com"
func Test017(t *testing.T) {
	main017()
}
