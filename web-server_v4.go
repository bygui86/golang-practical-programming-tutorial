package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
)

var washPostXML = []byte(`
<sitemapindex>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
   </sitemap>
</sitemapindex>
`)

type NewsAggPage struct {
	Title string
	News  map[string]NewsMap
}

type NewsMap struct {
	Keyword  string
	Location string
}

type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

func newsHandler(w http.ResponseWriter, r *http.Request) {
	var s Sitemapindex
	var n News
	// resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	// bytes, _ := ioutil.ReadAll(resp.Body)
	bytes := washPostXML
	xml.Unmarshal(bytes, &s)
	news_map := make(map[string]NewsMap)

	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		for idx, _ := range n.Keywords {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
		break // just first to avoid waste of time
	}

	p := NewsAggPage{Title: "Amazing News Aggregator", News: news_map}
	t, _ := template.ParseFiles("web-server_v4_templating.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", newsHandler)

	port := 8080
	fmt.Println("Starting server on port", port)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
