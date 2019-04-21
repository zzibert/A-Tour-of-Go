package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"../../cyoa"
)

func main() {
	filename := flag.String("file", "gopher.json", "The JSON file.")
	port := flag.Int("port", 3000, "The port to start the cyoa web application on")
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

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
