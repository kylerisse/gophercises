package main

import (
	"flag"
	"fmt"

	"./sitemap"
)

func main() {
	var (
		domain string
		depth  int
		delay  float64 //delay in seconds
	)
	flag.StringVar(&domain, "domain", "", "root domain to crawl")
	flag.IntVar(&depth, "depth", 4, "depth of crawling")
	flag.Float64Var(&delay, "delay", 0.5, "delay between http gets (in seconds)")
	flag.Parse()
	sitemap, err := sitemap.NewSitemap(domain)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for {
		if sitemap.Unvisited() < 1 || depth < 1 {
			break
		}
		sitemap.Crawl(delay)
		depth--
	}
	output, err := sitemap.ToXMLString()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	println(output)
}
