package urlshort

import (
	"database/sql"
	"encoding/json"
	"net/http"

	// postgres compatibility for sql
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

// PathURL has a short url and redirect path
type PathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if val, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, val, http.StatusSeeOther)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var items []PathURL
	err := yaml.Unmarshal(yml, &items)
	itemsMap := pathURLtoMap(items)
	return MapHandler(itemsMap, fallback), err
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
//
// JSON is expected to be in the format:
// [
// {"path": "/some-path", "url": "https://www.some-url.com/demo"},
// {"path": "/some-other-path", "url": "https://www.some-other-url.com/demo"}
// ]
func JSONHandler(jsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var items []PathURL
	err := json.Unmarshal(jsonBytes, &items)
	itemsMap := pathURLtoMap(items)
	return MapHandler(itemsMap, fallback), err
}

// DBHandler will get records from a database and then return
// an http.HandlerFunc (which also implements http.Handler)
//
// DB connection string should contain all relevant parameters
// a table named items is assumed
func DBHandler(connStr string, fallback http.Handler) (http.HandlerFunc, error) {
	itemMap := map[string]string{}
	if connStr == "" {
		return MapHandler(itemMap, fallback), nil
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return MapHandler(itemMap, fallback), err
	}
	rows, err := db.Query("SELECT * FROM items;")
	if err != nil {
		return MapHandler(itemMap, fallback), err
	}
	for rows.Next() {
		var row PathURL
		err := rows.Scan(&row.Path, &row.URL)
		if err != nil {
			continue
		}
		itemMap[row.Path] = row.URL
	}
	return MapHandler(itemMap, fallback), nil
}

func pathURLtoMap(pathURLs []PathURL) map[string]string {
	itemMap := map[string]string{}
	for i := range pathURLs {
		itemMap[pathURLs[i].Path] = pathURLs[i].URL
	}
	return itemMap
}
