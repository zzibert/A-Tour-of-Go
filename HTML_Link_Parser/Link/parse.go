package link

import "io"

//	Link respresent a Link in a HTML document.
type Link struct {
	Href string
	Text string
}

// Parse will take an HTML document and return a slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {

}
