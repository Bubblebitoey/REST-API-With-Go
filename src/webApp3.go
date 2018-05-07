package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)

	var s Sitemapindex
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		fmt.Printf("%s\n", Location)
	}
}