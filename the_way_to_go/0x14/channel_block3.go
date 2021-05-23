package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // 创建通道
	go func() {
		time.Sleep(15 * 1e9)
		x := <-c // 接收
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	c <- 10 // 发送
	fmt.Println("sent", 10)
}
