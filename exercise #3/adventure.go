package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func NewsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amazing news Aggregator", News: "trololo trololo trololo"}
	t, _ := template.ParseFiles("basicTemplating.html")
	t.Execute(w, p)
}

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func main() {
	file, _ := ioutil.ReadFile("./gopher.json")
	chapters := make(map[string]Chapter)

	err := json.Unmarshal(file, &chapters)

	if err != nil {
		fmt.Println(err)
	}

	for key, value := range chapters {
		fmt.Printf("%s is for %s\n", key, value.Title)

	}

	http.HandleFunc("/", introHandler)
	http.HandleFunc("/new-york", newYorkHandler)
	http.HandleFunc("/debate", debatekHandler)
	http.HandleFunc("/sean-kelly", seanKellyHandler)
	http.HandleFunc("/mark-bates", markBatesHandler)
	http.HandleFunc("/denver", denverHandler)
	http.HandleFunc("/home", homeHandler)

	http.ListenAndServe(":8000", nil)
}
