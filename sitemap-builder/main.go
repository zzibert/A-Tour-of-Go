package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	link "github.com/zzibert/A-Tour-of-Go/HTML_Link_Parser/Link"
)

func main() {
	//	Creating the url flag
	urlPtr := flag.String("url", "https://gophercises.com/", "The website url")
	depthPtr := flag.Int("depth", 3, "The depth of the bfs")
	flag.Parse()

	queue := make([]link.Link, 0)

	Bfs(*urlPtr, &queue, *depthPtr)

	fmt.Printf("%+v\n", queue)

}

func filterLinks(domain string, links []link.Link) []link.Link {
	newLinks := make([]link.Link, 0)
	for _, link := range links {
		if strings.Contains(link.Href, domain) {
			newLinks = append(newLinks, link)
		}
	}
	return newLinks
}

// Bfs stands for breath-first-search.
func Bfs(urlString string, queue *[]link.Link, depth int) {
	if depth < 1 {
		return
	}

	response, err := http.Get(urlString)
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	r := bytes.NewReader(body)
	u, _ := url.Parse(urlString)
	links, err := link.Parse(r, u)
	if err != nil {
		fmt.Println(err)
	}
	links = filterLinks(u.String(), links)
	for _, link := range links {
		if !checkIfInqueue(link, queue) {
			*queue = append(*queue, link)
			Bfs(link.Href, queue, depth-1)
		}
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
