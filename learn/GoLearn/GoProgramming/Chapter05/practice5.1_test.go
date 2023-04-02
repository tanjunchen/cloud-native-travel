package Chapter01

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"testing"
)

// go test -v -run TestStart practice5.1_test.go -args https://www.baidu.com
func TestStart(t *testing.T) {
	Start()
}

func Start() {
	reps, err := http.Get(os.Args[4])
	if err != nil {
		fmt.Println(err)
	}
	doc, err := html.Parse(io.Reader(reps.Body))
	fmt.Println(doc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlink: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links [] string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

/**
visit(visit(links, n.FirstChild), n.NextSibling)
函数如上所述
*/
