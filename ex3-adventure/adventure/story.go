package adventure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// Story is a map of Chapters with an identifier
// the "intro" key is the first chapter
type Story map[string]Chapter

// Chapter has a title, multiple paragraphs and
// multiple options
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Choices    []Choice `json:"options"`
}

// Choice defines the next Chapter and teaser text
type Choice struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

// ChapterHandler method returns the web page with current chapter
func (s Story) ChapterHandler(chapter string, chTemplate string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(chTemplate)
		if err != nil {
			fmt.Println("cannot parse template file\n" + err.Error())
			os.Exit(1)
		}
		err = tmpl.ExecuteTemplate(w, "chapter.html", s[chapter])
	}
}

// CreateStory from a source JSON file
func CreateStory(fileName string) (Story, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	buf.ReadFrom(f)
	var s Story
	err = json.Unmarshal(buf.Bytes(), &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
