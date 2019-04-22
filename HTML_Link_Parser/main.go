package main

import (
	"fmt"

	"./link"
)

func main() {
	link := link.Link{"Href", "Text"}
	fmt.Printf("LOL %s", link.Text)
}
