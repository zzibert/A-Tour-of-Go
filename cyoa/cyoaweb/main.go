package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"../../cyoa"
)

func main() {
	filename := flag.String("file", "gopher.json", "The JSON file.")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		fmt.Println(err)
	}
	d := json.NewDecoder(f)
	var story cyoa.Story
	if err := d.Decode(&story); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", story)
}
