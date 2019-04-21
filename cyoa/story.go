package cyoa

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

// func init() {
// 	tpl := template.Must(template.New("").Parse(defaultHandlerTemplate))
// }

var defaultHandlerTemplate = `<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Choose Your Own Adventure</title>
</head>
<body>
    <h1>{{ .Title }}</h1>
    {{ range .Paragraphs }}
        <p>{{ . }}</p>
    {{ end }}
    <ul>
        {{ range .Options }}
        <li><a href="/{{ .Chapter }}">{{ .Text }}</a></li>
        {{ end }}
    </ul>
</body>
</html>`

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("").Parse(defaultHandlerTemplate))
	path := r.URL.Path
	if path == "/" {
		path = "intro"
	} else {
		path = path[1:]
	}
	err := tpl.Execute(w, h.s[path])
	if err != nil {
		fmt.Println(err)
	}
}

func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

type Story map[string]Chapter
