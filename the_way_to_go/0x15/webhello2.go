package main

import (
	"fmt"
	"net/http"
	"strings"
)

// http://localhost:8080/hello/jck
func helloHandler(w http.ResponseWriter, r *http.Request) {
	remPartOfURL := r.URL.Path[len("/hello/"):] //get everything after the /hello/ part of the URL
	fmt.Fprintf(w, "Hello %s!", remPartOfURL)
}

// http://localhost:8080/shouthello/jck
func shouthelloHandler(w http.ResponseWriter, r *http.Request) {
	remPartOfURL := r.URL.Path[len("/shouthello/"):] //get everything after the /shouthello/ part of the URL
	fmt.Fprintf(w, "Hello %s!", strings.ToUpper(remPartOfURL))
}

// go run webhello2.go
func main() {
	http.HandleFunc("/hello/", helloHandler)
	http.HandleFunc("/shouthello/", shouthelloHandler)
	http.ListenAndServe("localhost:8080", nil)
}
