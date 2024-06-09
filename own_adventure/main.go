package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type story struct {
	Title   string        `json:"title"`
	Story   []string      `json:"story"`
	Options []storyOption `json:"options"`
}

type storyOption struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Load struct{}

type StoryData struct {
	PageTitle string
	Title     string
	Stories   []string
	Options   []storyOption
}

func main() {
	var load Load
	http.Handle("/", load)
	http.ListenAndServe(":8080", nil)
}

func (Load) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	p = strings.ReplaceAll(p, "/", "")
	if p == "" {
		p = "intro"
	}

	d, err := readJson("gopher.json")
	if err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseFiles("layout.html"))
	data := StoryData{
		PageTitle: "Make Your Own Adventure",
		Title:     d[p].Title,
		Stories:   d[p].Story,
		Options:   d[p].Options,
	}
	tmpl.Execute(w, data)
}

func readJson(fileName string) (map[string]story, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	var jData map[string]story
	dec := json.NewDecoder(f)
	if err := dec.Decode(&jData); err != nil {
		return nil, err
	}
	return jData, nil
}
