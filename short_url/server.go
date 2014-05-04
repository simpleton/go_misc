package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It Works!!! %s", r.URL.Path[1:])
}

func handler_redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://simsun.me", 301)
}

func main() {
	http.HandleFunc("/", handler_redirect)
	http.ListenAndServe(":8080", nil)
}


