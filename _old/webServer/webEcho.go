package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var counter int
var mu sync.Mutex

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/counter/", counterHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request from %v: %s\n", r.RemoteAddr, r.URL.RequestURI())

	mu.Lock()
	counter++
	mu.Unlock()

}

func counterHandler(w http.ResponseWriter, r *http.Request) {

	mu.Lock()
	fmt.Fprintf(w, "Query Count: %d\n", counter)
	mu.Unlock()

}
