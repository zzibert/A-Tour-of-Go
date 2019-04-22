package main

import (
	"fmt"
	"strings"

	"./Link"
)

var exampleHTML = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/dog">
  <span>Something in a span</span>
  Text not in a span
  <b>Bold text!</b>
</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHTML)
	links, err := link.Parse(r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", links)
}
