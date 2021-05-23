package main

import (
	"fmt"
)

func f1(in chan int) {
	fmt.Println(<-in) // 阻塞
}

func main() {
	out := make(chan int) // 创建通道
	//out := make(chan int, 1) // solution 2
	//go f1(out)  // solution 1
	out <- 2 // 阻塞
	go f1(out)
}
