package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
	fmt.Fprintf(w, "<p>Go is fast!</p>")
	fmt.Fprintf(w, "<p>...and simple!</p>")
	fmt.Fprintf(w, "<p>You %s even add %s</p>", "can", "<strong>variables</strong>")
	fmt.Fprintf(w, `
		<h6>You can even do ...</h6>
		<h5>multiple lines ...</h5>
		<h4>in one %s</h4>`, "formatted print")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Greeting from Rabbit :)")
}

func main() {
	http.HandleFunc("/", index_handler)       // path, then what function to run
	http.HandleFunc("/about/", about_handler) // path, then what function to run

	port := 8080
	fmt.Println("Starting server on port", port)
	http.ListenAndServe(":"+strconv.Itoa(port), nil) // listen on what port?   ... can serve on tls with ListenAndServeTLS ... need something in server args, we'll put nil for now
}
