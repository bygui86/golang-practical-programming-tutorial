package main

import (
	"encoding/xml"
	"fmt"
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

type Sitemapindex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

// Each element in this Go slice is {url} rather than just straight up url. Why is this? It's fairly subtle,
// but each of these URLs is actually a Location type, which just so happens to have a Loc element that is the string
// we're looking for. We'd rather this slice that we're building to be a slice of strings, not a slice of Location types.
// To fix this, we need to build a string method to associate with our Location type.
// PLEASE NOTE: the method starts with CAPITOL LETTER, this was can be used without explicit invocation (see line 50)
func (el Location) String() string {
	// We're using fmt.Sprintf() here, which takes in data, formats according to a specifier, and returns the resulting string.
	// More information can be found at Sprintf in the Golang Documentation.
	return fmt.Sprintf(l.Loc)
}

func main() {
	// The inevitable reality of Washington Post changing their structure ...
	// resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	// bytes, _ := ioutil.ReadAll(resp.Body)
	// ... but we can simulate the same structure
	bytes := washPostXML
	var s Sitemapindex
	xml.Unmarshal(bytes, &s)
	fmt.Println(s.Locations)
}
