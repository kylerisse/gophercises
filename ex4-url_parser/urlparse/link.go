package urlparse

import (
	"bufio"
	"strings"

	"golang.org/x/net/html"
)

// Link contains a reference and the assoiciated text
type Link struct {
	Href string
	Text string
}

// ExtractLinks returns a slice of Links extracted from HTML
func ExtractLinks(sr *bufio.Reader) ([]Link, error) {
	doc, err := html.Parse(sr)
	if err != nil {
		return nil, err
	}
	var links []Link
	var f func(n *html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, Link{
						Href: a.Val,
						Text: cleanupText(n.FirstChild.Data),
					})
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return links, nil
}

func cleanupText(s string) string {
	s = strings.Split(s, "<!---")[0]
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\t", "")
	s = strings.Trim(s, " ")
	return s
}
