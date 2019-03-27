package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
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

// Note the use of the '>' to specify embedded tags.
type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

func getSitemap(bytes []byte, print bool) Sitemapindex {
	var s Sitemapindex
	xml.Unmarshal(bytes, &s)
	if print {
		fmt.Println("Locations:")
		for _, Location := range s.Locations {
			fmt.Printf("%s\n", Location)
		}
	}
	return s
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

func getNews(s Sitemapindex) map[string]NewsMap {
	var n News
	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		break // just first to avoid waste time
	}
	fmt.Printf("News:")
	fmt.Printf("\t%d Titles\n", len(n.Titles))
	fmt.Printf("\t%d Keywords\n", len(n.Keywords))
	fmt.Printf("\t%d Locations\n", len(n.Locations))

	news_map := make(map[string]NewsMap)
	for idx, _ := range n.Titles {
		news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
	}
	return news_map
}

func printMap(news_map map[string]NewsMap) {
	fmt.Printf("News map:")
	for idx, data := range news_map {
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)
	}
}

func main() {
	var print bool = true
	s := getSitemap(washPostXML, print)
	fmt.Println()
	nm := getNews(s)
	fmt.Println()
	printMap(nm)
}
