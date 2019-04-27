package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	link "github.com/zzibert/A-Tour-of-Go/HTML_Link_Parser/Link"
)

func main() {
	//	Creating the url flag
	urlPtr := flag.String("url", "https://en.wikipedia.org/wiki/Main_Page", "The website url")
	flag.Parse()

	response, err := http.Get(*urlPtr)
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	r := bytes.NewReader(body)
	links, err := link.Parse(r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", links)

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

func bfs() {

}
