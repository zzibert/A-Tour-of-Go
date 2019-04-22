package link

import (
	"io"
	"strings"

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
	return dfs(doc), nil
}

func getText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	text := ""
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text = text + getText(c)
	}
	return strings.Join(strings.Fields(text), " ")
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}

func buildLink(n *html.Node) Link {
	newLink := Link{"", ""}
	for _, a := range n.Attr {
		if a.Key == "href" {
			newLink.Href = a.Val
			break
		}
	}
	newLink.Text = getText(n)
	return newLink
}

func dfs(n *html.Node) []Link {
	nodes := linkNodes(n)
	links := make([]Link, 0)
	for _, link := range nodes {
		links = append(links, buildLink(link))
	}
	return links
}
