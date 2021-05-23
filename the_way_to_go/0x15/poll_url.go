package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"http://www.bing.com/",
	"http://python.org/",
}

func main() {
	// Execute an HTTP HEAD request for all url's
	// and returns the HTTP status string or an error string.
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error", url, err)
		}
		fmt.Println(url, ": ", resp.Status)
	}
}

/* Output:
http://www.bing.com/ :  200 OK
http://python.org/ :  200 OK
*/
