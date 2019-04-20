package main

import (
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
	story, err := cyoa.JsonStory(f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", story)
}
