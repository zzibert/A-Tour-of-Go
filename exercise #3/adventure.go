package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

// func introHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
// }

func introHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "trolollo", Paragraphs: []string{"ena", "dva", "tri"}}
	t, _ := template.ParseFiles("gopher.html")
	t.Execute(w, p)
}

func Handler(url string, page Page) http.Handler {
	return http.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("gopher.html")
		t.Execute(w, page)
	})
}

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

type Page struct {
	Title      string
	Paragraphs []string
}

func main() {
	file, _ := ioutil.ReadFile("./gopher.json")
	chapters := make(map[string]Chapter)
	pages := make(map[string]Page)
	err := json.Unmarshal(file, &chapters)

	if err != nil {
		fmt.Println(err)
	}

	for key, value := range chapters {
		pages[key] = Page{Title: value.Title, Paragraphs: value.Paragraphs}
	}

	http.HandleFunc("/", introHandler)
	// http.HandleFunc("/new-york", newYorkHandler)
	// http.HandleFunc("/debate", debatekHandler)
	// http.HandleFunc("/sean-kelly", seanKellyHandler)
	// http.HandleFunc("/mark-bates", markBatesHandler)
	// http.HandleFunc("/denver", denverHandler)
	// http.HandleFunc("/home", homeHandler)

	http.ListenAndServe(":8000", nil)
}
