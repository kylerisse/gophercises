package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"

	"./urlshort"
)

func main() {
	var yamlFile string
	flag.StringVar(&yamlFile, "y", "", "specify yaml file to load")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yaml := readYamlFile(yamlFile)

	// YAMLHandler using the mapHandler as the fallback
	yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func readYamlFile(fileName string) []byte {

	backupYaml := []byte(`
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`)
	if fileName == "" {
		return backupYaml
	}

	f, err := os.Open(fileName)
	if err != nil {
		return backupYaml
	}

	var fileYaml bytes.Buffer
	fileYaml.ReadFrom(f)

	return fileYaml.Bytes()
}
