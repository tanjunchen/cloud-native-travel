package Chapter01

import (
	"golang.org/x/net/html"
	"testing"
)

func TestStart058(t *testing.T) {

}

//func ElementByID(doc *html.Node, id string) *html.Node {
//	var elem *html.Node
//	forEachNode(doc,
//		func(n *html.Node) bool {
//			if n.Type == html.ElementNode {
//				for _, a := range n.Attr {
//					if a.Key == "id" && a.Val == id {
//						elem = n
//						return false
//					}
//				}
//			}
//			return true
//		},
//		nil)
//	return elem
//}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
