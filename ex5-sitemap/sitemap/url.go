package sitemap

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func (s *Sitemap) fetchURL(url string) ([]byte, error) {
	resp, err := http.Get("https://" + s.domain + url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
}

func extractRefs(bys []byte) []string {
	doc, err := html.Parse(bytes.NewReader(bys))
	if err != nil {
		return []string{""}
	}
	var urls []string
	var f func(n *html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					urls = append(urls, a.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return urls
}

func (s *Sitemap) normalizePath(url string) string {
	var path string
	if strings.HasPrefix(url, "/") && !strings.HasPrefix(url, "//") {
		path = url
	}
	if strings.Contains(url, s.domain) {
		path = strings.Split(url, s.domain)[1]
	}
	if strings.HasSuffix(path, "/") {
		return path
	}
	return path + "/"
}

func (s *Sitemap) inDomain(url string) bool {
	if strings.HasPrefix(url, "/") {
		return true
	}
	if strings.HasPrefix(url, "https://"+s.domain) {
		return true
	}
	if strings.HasPrefix(url, "http://"+s.domain) {
		return true
	}
	return false
}
