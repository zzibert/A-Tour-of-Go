package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type NewsAggPage struct {
	Title string
	News  string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func NewsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amazing news Aggregator", News: "trololo trololo trololo"}
	t, _ := template.ParseFiles("basicTemplating.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", NewsAggHandler)
	http.ListenAndServe(":8000", nil)
}
