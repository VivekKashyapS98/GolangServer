package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Server has started, current route is %v </h1>", r.URL.Path[0:])
	fmt.Printf("Server has started , current route is %v", r.URL.Path[0:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
