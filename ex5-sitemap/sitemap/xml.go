package sitemap

import (
	"encoding/xml"
	"sort"
)

// XMLNS is default sitemap schema url
const XMLNS string = "http://www.sitemaps.org/schemas/sitemap/0.9"

// XMLSitemap for creating XML output
type XMLSitemap struct {
	XMLName string `xml:"urlset"`
	Xmlns   string `xml:"xmlns,attr"`
	URLs    []URL  `xml:"url"`
}

// URL for XMLSitemap
type URL struct {
	Loc string `xml:"loc"`
}

// ByAlpha to provide sort interface to []URL
type ByAlpha []URL

// ToXMLString Siteemap in XML format
func (s *Sitemap) ToXMLString() (string, error) {
	x := XMLSitemap{
		XMLName: "sitemap",
		Xmlns:   XMLNS,
		URLs:    []URL{},
	}
	for k := range s.urls {
		x.URLs = append(x.URLs, URL{
			Loc: "https://" + s.domain + k,
		})
	}
	sort.Sort(ByAlpha(x.URLs))
	buf, err := xml.MarshalIndent(x, "", "  ")
	if err != nil {
		return "", err
	}
	return xml.Header + string(buf), nil
}

// Len of []URL
func (b ByAlpha) Len() int {
	return len(b)
}

// Less of []URL
func (b ByAlpha) Less(i, j int) bool {
	return b[i].Loc < b[j].Loc
}

// Swap of []URL
func (b ByAlpha) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
