package Chapter01

import (
	"golang.org/x/net/html"
	"testing"
)

func TestStart017(t *testing.T) {

}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	if len(name) == 0 {
		return nil
	}
	var visit func(elems []*html.Node, n *html.Node) []*html.Node
	visit = func(elems []*html.Node, n *html.Node) []*html.Node {
		for _, tag := range name {
			if n.Type == html.ElementNode && n.Data == tag {
				elems = append(elems, n)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			elems = visit(elems, c)
		}
		return elems
	}
	return visit(nil, doc)
}