package builder

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	link "github.com/zzibert/A-Tour-of-Go/HTML_Link_Parser/Link"
)

func FilterLinks(domain string, links []link.Link) []link.Link {
	newLinks := make([]link.Link, 0)
	for _, link := range links {
		if strings.Contains(link.Href, domain) {
			newLinks = append(newLinks, link)
		}
	}
	return newLinks
}

func Bfs(url string, queue *[]link.Link) {

	response, err := http.Get(url)
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
	links = FilterLinks(url, links)
	for _, link := range links {
		if !checkIfInqueue(link, queue) {
			*queue = append(*queue, link)
			Bfs(link.Href, queue)
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
