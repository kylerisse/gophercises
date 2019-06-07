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
	var jsonFile string
	flag.StringVar(&jsonFile, "j", "", "specify json file to load")
	var dbString string
	flag.StringVar(&dbString, "d", "postgresql://postgres:password@localhost/postgres?sslmode=disable", "specify a DB connection string")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yaml := readYAMLfile(yamlFile)

	// YAMLHandler using the mapHandler as the fallback
	yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
	if err != nil {
		fmt.Println(err.Error())
	}

	jsonBytes := readJSONfile(jsonFile)

	// JSONHandler using the yamlHandler as the fallback
	jsonHandler, err := urlshort.JSONHandler(jsonBytes, yamlHandler)
	if err != nil {
		fmt.Println(err.Error())
	}

	// DBHandler using the jsonHandler as the fallback
	dbHandler, err := urlshort.DBHandler(dbString, jsonHandler)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", dbHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func fileToBytes(fileName string, fallback []byte) []byte {
	if fileName == "" {
		return fallback
	}
	f, err := os.Open(fileName)
	if err != nil {
		return fallback
	}
	var fileContent bytes.Buffer
	fileContent.ReadFrom(f)

	return fileContent.Bytes()
}

func readYAMLfile(fileName string) []byte {
	backupYaml := []byte(`
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`)
	return fileToBytes(fileName, backupYaml)
}

func readJSONfile(fileName string) []byte {
	backupJSON := []byte(`
[
	{"path": "/gc", "url": "https://greatercommons.com/"},
	{"path": "/techlead", "url": "http://youtube.com/techlead"}
]
`)
	return fileToBytes(fileName, backupJSON)
}
