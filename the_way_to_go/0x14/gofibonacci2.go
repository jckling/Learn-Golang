package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // 关闭通道
}

func main() {
	c := make(chan int, 25) // 缓冲
	start := time.Now()
	go fibonacci(cap(c), c) // 容量
	for i := range c {
		fmt.Println(i)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
