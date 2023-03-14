// Ref: Section 1.7, Chapter 1
// Alan A. A. Donovan and Brian W. Kernighan, The Go Programming Langugage
package server

import (
	"fmt"
	"log"
	"net/http"
)

func server() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
