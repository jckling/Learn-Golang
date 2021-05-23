package main

import (
	"flag"
	"fmt"
)

var ngoroutine = flag.Int("n", 100000, "how many goroutines") // 默认值 100000

func f(left, right chan int) { left <- 1 + <-right }

func main() {
	flag.Parse()
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost
	for i := 0; i < *ngoroutine; i++ {
		left, right = right, make(chan int)
		go f(left, right) // 阻塞
	}
	right <- 0      // bang!
	x := <-leftmost // wait for completion
	fmt.Println(x)  // 100000, about 1.5 s
}
