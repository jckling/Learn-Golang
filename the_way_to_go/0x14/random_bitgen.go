package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// consumer:
	go func() {
		for {
			fmt.Print(<-c, " ") // ζ ι 01 εΊε
		}
	}()

	// producer:
	for {
		select {
		case c <- 0:
		case c <- 1:
		}
	}

}
