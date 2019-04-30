package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	link "github.com/zzibert/A-Tour-of-Go/HTML_Link_Parser/Link"
)

type loc struct {
	loc string `xml:loc`
}

type urlset struct {
	url []loc `xml:url`
}

func main() {
	//	Creating the url flag
	urlPtr := flag.String("url", "https://gophercises.com/", "The website url")
	depthPtr := flag.Int("depth", 3, "The depth of the bfs")
	flag.Parse()

	seen := make(map[string]bool)

	Bfs(*urlPtr, seen, *depthPtr)

	// var toXml urlset
	// for _, page := range queue {
	// 	toXml.url = append(toXml.url, loc{page.Href})
	// }

	// encoder := xml.NewEncoder(os.Stdout)
	// encoder.Indent("", "  ")
	// if err := encoder.Encode(toXml); err != nil {
	// 	fmt.Println(err)
	// }

	for k, _ := range seen {
		fmt.Println(k)
	}

}

func filterLinks(domain string, links []link.Link) []link.Link {
	newLinks := make([]link.Link, 0)
	for _, link := range links {
		if strings.HasPrefix(link.Href, domain) {
			newLinks = append(newLinks, link)
		}
	}
	return newLinks
}

// Bfs stands for breath-first-search.
func Bfs(urlString string, seen map[string]bool, depth int) {
	if depth < 1 {
		return
	}
	for _, l := range get(urlString) {
		if seen[l.Href] {
			continue
		}
		seen[l.Href] = true
		Bfs(l.Href, seen, depth-1)
		// if !checkIfInqueue(l, queue) {
		// 	*queue = append(*queue, l)
		// 	Bfs(l.Href, queue, depth-1)
		// }
	}
}

func checkIfInqueue(link link.Link, queue *[]link.Link) bool {
	for _, queueLink := range *queue {
		if link.Href == queueLink.Href {
			return true
		}
	}
	return false
}

func get(urlString string) []link.Link {
	response, err := http.Get(urlString)
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	u, _ := url.Parse(urlString)
	links, err := link.Parse(response.Body, u)
	if err != nil {
		fmt.Println(err)
	}
	links = filterLinks(u.String(), links)
	return links
}
