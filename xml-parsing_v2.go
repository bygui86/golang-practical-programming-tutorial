package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// In Go we refer to an array as a 'slice'
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

func getNews(s Sitemapindex) {
	var n News
	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		break
	}
	fmt.Printf("Found %d Titles\n", len(n.Titles))
	fmt.Printf("Found %d Keywords\n", len(n.Keywords))
	fmt.Printf("Found %d Locations\n", len(n.Locations))
}

func main() {
	var print bool = false
	s := getSitemap(washPostXML, print)
	getNews(s)
}
