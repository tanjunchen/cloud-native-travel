package Chapter01

import (
	"golang.org/x/net/html"
)

var filter = map[string]string{
	"a":      "href",
	"img":    "src",
	"script": "src",
}
// 用一个字典匹配所有的类型，先匹配n.Data，在遍历属性得到需要的地址
func visit4(links [] string, n *html.Node) []string {
	for k, v := range filter {
		if n.Type == html.ElementNode && n.Data == k {
			for _, a := range n.Attr {
				if a.Key == v {
					links = append(links, a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit4(links, c)
	}
	return links
}
