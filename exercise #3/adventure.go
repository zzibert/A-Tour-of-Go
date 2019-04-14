package main

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
// }

// func NewsAggHandler(w http.ResponseWriter, r *http.Request) {
// 	p := NewsAggPage{Title: "Amazing news Aggregator", News: "trololo trololo trololo"}
// 	t, _ := template.ParseFiles("basicTemplating.html")
// 	t.Execute(w, p)
// }
type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func main() {
	file, _ := ioutil.ReadFile("./gopher.json")
	chapters := new(map[string]Chapter)

	if err := json.Unmarshal(file, &chapters)

	// for _, chapter := range chapters {
	// 	fmt.Println(chapter.Title)
	// }

	// http.HandleFunc("/", indexHandler)
	// http.HandleFunc("/agg/", NewsAggHandler)
	// http.ListenAndServe(":8000", nil)
}
