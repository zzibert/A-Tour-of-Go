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
func bfs(urlString string, maxDepth int) []string {
	seen := make(map[string]struct{})
	// var q map[string]struct{}
	nq := map[string]struct{}{
		urlString: struct{}{},
	}

	for i := 0; i < maxDepth; i++ {
		q, nq := nq, make(map[string]struct{})
		for url := range q {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = struct{}{}
			for _, link := range get(url) {
				nq[link.Href] = struct{}{}
			}
		}
	}
	ret := make([]string, 0)
	for key := range seen {
		ret = append(ret, key)
	}
	return ret
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
