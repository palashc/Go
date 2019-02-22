package cyoa

import (
	"encoding/json"
	"io"
	"net/http"
	"text/template"
)

func JsonStory(r io.Reader) (Story, error) {
	json_decoder := json.NewDecoder(r)
	var story Story
	if err := json_decoder.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("").Parse(defaultHTTPTemplate))
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
}

var defaultHTTPTemplate = `
							<!DOCTYPE html>
							<html>
							<head>
								<meta charset="utf-8">
								<title>Choose your own adventure!</title>
							</head>
							<body>
								<h1>{{.Title}}</h1>
								{{range .Paragraphs}}
								<p>{{.}}</p>
								{{end}}
								<ul>
									{{range .Options}}
									<li><a href="/{{.Chapter}}">{{.Text}}</a></li>
									{{end}}
								</ul>
							</body>
							</html>
							`
