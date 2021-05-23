package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.bing.com")
	CheckError(err)

	data, err := ioutil.ReadAll(res.Body) // 网页内容
	CheckError(err)

	fmt.Printf("Got: %q", string(data))
}

func CheckError(err error) {
	if err != nil {
		log.Fatalf("Get: %v", err)
	}
}
