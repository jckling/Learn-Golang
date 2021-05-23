package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 50) // 缓冲区
	go func() {
		time.Sleep(15 * 1e9)
		x := <-c
		fmt.Println("received", x) // 不执行
	}()
	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)
}
