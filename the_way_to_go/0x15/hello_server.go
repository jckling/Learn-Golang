package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

// go run hello_server.go
func main() {
	var h Hello
	http.ListenAndServe("localhost:8080", h)
}
