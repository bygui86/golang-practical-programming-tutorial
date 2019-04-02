package main

import (
	"fmt"
	"html/template" // WARNING: specify the right import becuase standard import could be 'text/template'
	"net/http"
	"strconv"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

type NewsAggPage struct {
	Title string
	News  string
}

func agg_handler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amazing News Aggregator", News: "... some news ..."}
	t, _ := template.ParseFiles("web-server_v3_templating.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/agg/", agg_handler)

	port := 8080
	fmt.Println("Starting server on port", port)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
