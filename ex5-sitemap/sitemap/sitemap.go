package sitemap

import (
	"net"
	"time"
)

// Sitemap for a domain contains URLs
type Sitemap struct {
	domain string
	urls   map[string]bool
}

// NewSitemap constructor
func NewSitemap(domain string) (Sitemap, error) {
	_, err := net.LookupHost(domain)
	if err != nil {
		return Sitemap{}, err
	}
	return Sitemap{
		domain: domain,
		urls: map[string]bool{
			"/": false,
		},
	}, nil
}

// Unvisited count of unvisited URLs
func (s *Sitemap) Unvisited() int {
	var count int
	for _, v := range s.urls {
		if !v {
			count++
		}
	}
	return count
}

// Crawl Unvisited URLs
func (s *Sitemap) Crawl(delay float64) {
	var todo []string
	for k, v := range s.urls {
		if !v {
			b, err := s.fetchURL(k)
			time.Sleep(time.Duration(delay) + time.Second)
			s.urls[k] = true
			if err != nil {
				continue
			}
			for _, v := range extractRefs(b) {
				if !s.inDomain(v) {
					continue
				}
				todo = append(todo, v)
			}
		}
	}
	for _, v := range todo {
		path := s.normalizePath(v)
		_, ok := s.urls[path]
		if !ok {
			s.urls[path] = false
		}
	}
}
