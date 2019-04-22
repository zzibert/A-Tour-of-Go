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
	links := make([]Link, 0)
	if err != nil {
		return nil, err
	}
	dfs(doc, &links)
	return links, nil
}

func getText(newLink *Link, n *html.Node) {
	if n.Type == html.TextNode {
		newLink.Text = newLink.Text + strings.TrimSpace(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			newLink.Text = newLink.Text + strings.TrimSpace(c.Data)
		} else {
			getText(newLink, c)
		}
	}
}

func dfs(n *html.Node, links *[]Link) {
	if n.Type == html.ElementNode && n.Data == "a" {
		newLink := Link{"", ""}
		for _, a := range n.Attr {
			if a.Key == "href" {
				newLink.Href = a.Val
				break
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			getText(&newLink, c)
		}
		*links = append(*links, newLink)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, links)
	}
}
