package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// consumer:
	go func() {
		for {
			fmt.Print(<-c, " ") // 无限 01 序列
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
