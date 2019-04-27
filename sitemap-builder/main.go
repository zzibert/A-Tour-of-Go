package main

import (
	"flag"
	"fmt"

	link "github.com/zzibert/A-Tour-of-Go/HTML_Link_Parser/Link"

	builder "github.com/zzibert/A-Tour-of-Go/sitemap-builder/builder"
)

func main() {
	//	Creating the url flag
	urlPtr := flag.String("url", "https://www.comtrade.com/", "The website url")
	flag.Parse()

	queue := make([]link.Link, 0)

	builder.bfs(*urlPtr, &queue)

	fmt.Printf("%+v\n", queue)

}
