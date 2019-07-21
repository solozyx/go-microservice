package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":10000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Web")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("health check")
}
