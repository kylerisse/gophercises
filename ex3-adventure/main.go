package main

import (
	"fmt"
	"net/http"

	"./adventure"
)

func main() {
	chTemplate := "./templates/chapter.html"
	story, err := adventure.CreateStory("gopher.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	mux := setMux(story, chTemplate)
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mux)
}

func setMux(st adventure.Story, chTemplate string) *http.ServeMux {
	mux := http.NewServeMux()
	for k := range st {
		fmt.Println("Adding chapter " + k + " to story.")
		mux.HandleFunc("/chapter/"+k, st.ChapterHandler(k, chTemplate))
	}
	mux.HandleFunc("/", st.ChapterHandler("intro", chTemplate))
	return mux
}
