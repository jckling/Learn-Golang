package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	go pump(ch1)
	fmt.Println(<-ch1) // 从通道中提取一个数据，通道可用
}

// 生产者
func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i // 发送一个数据到通道，阻塞
	}
}
