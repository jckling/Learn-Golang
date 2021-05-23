package main

import (
	"fmt"
)

func sum(x, y int, c chan int) {
	c <- x + y
}

func main() {
	c := make(chan int)
	go sum(12, 13, c) // 通道求值
	fmt.Println(<-c)  // 25
}
