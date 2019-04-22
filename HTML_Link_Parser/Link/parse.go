package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

//	Link respresent a Link in a HTML document.
type Link struct {
	Href string
	Text string
}

// Parse will take an HTML document and return a slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	dfs(doc)
	return nil, nil
}

func dfs(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode {
				fmt.Println(c.Data)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c)
	}
}
