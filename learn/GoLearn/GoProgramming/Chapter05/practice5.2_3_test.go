package Chapter01

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"testing"
)

func TestStart052(t *testing.T) {
	Start052()
}

func Start052() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findelems: %v\n", err)
		os.Exit(1)
	}
	elements := map[string]int{}
	visit2(elements, doc)
	for elem, count := range elements {
		fmt.Printf("%s\t%d\n", elem, count)
	}
}

func visit2(e map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		e[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit2(e, c)
	}
}

func Start053() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findtexts: %v\n", err)
		os.Exit(1)
	}
	visit3(doc)
}

func TestStart053(t *testing.T) {
	Start053()
}

func visit3(n *html.Node) {
	if n != nil && n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" {
			return
		}
	}
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit3(c)
	}
}
