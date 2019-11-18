package main

import (
	"encoding/xml"
	"fmt"
)

func compact() {
	fmt.Println("Compact:")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func external_var() {
	fmt.Println("External variable:")
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func infinite() {
	fmt.Println("Infinite loop:")
	for {
		fmt.Println("Do stuff.")
	}
}

func break_loop() {
	fmt.Println("Loop with break:")
	x := 5
	for {
		fmt.Println("Do stuff.", x)
		x += 3
		if x > 25 {
			break
		}
	}
}

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

func (e Location) String() string {
	return fmt.Sprintf(e.Loc)
}

// We want to loop over a list in a more traditional sense. We can do this by using Go's range, iterate over elements in data structures.
func range_loop() {
	bytes := washPostXML
	var s Sitemapindex
	xml.Unmarshal(bytes, &s)

	fmt.Println("Range loop:")
	// The first value returned will be the index of the value, which we're designating to _, since we aren't going to be using index for any reason.
	// If we don't use the underscore here, and instead use a variable name like index or something, Go will throw an error because we're not using the variable we defined.
	for _, Location := range s.Locations {
		fmt.Printf("%s\n", Location)
	}
}

func main() {
	compact()
	fmt.Println()
	external_var()
	fmt.Println()
	break_loop()
	fmt.Println()
	range_loop()
}
