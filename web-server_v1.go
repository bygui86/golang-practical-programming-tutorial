package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the INDEX page")
	// fmt.Fprintf(w, "This is the INDEX page", r)
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the ABOUT page")
}

func main() {
	http.HandleFunc("/", index_handler)       // path, then what function to run
	http.HandleFunc("/about/", about_handler) // path, then what function to run

	port := 8080
	fmt.Println("Starting server on port", port)
	http.ListenAndServe(":"+strconv.Itoa(port), nil) // listen on what port?   ... can serve on tls with ListenAndServeTLS ... need something in server args, we'll put nil for now
}
