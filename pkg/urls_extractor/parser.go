package urlsextractor

import (
	"io"

	"golang.org/x/net/html"
)

func Extract(body io.Reader) ([]string, error) {
	doc, err := html.Parse(body)
	if err != nil {
		return nil, err
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				links = append(links, a.Val)
			}
		}
	}
	forEachNode(doc, visitNode)
	return links, nil
}

func forEachNode(doc *html.Node, f func(n *html.Node)) {
	if f == nil {
		return
	}
	for node := range doc.Descendants() {
		f(node)
	}
}
