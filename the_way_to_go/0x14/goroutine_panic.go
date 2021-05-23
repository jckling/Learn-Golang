package main

import (
	"fmt"
)

// 生产者
func tel(ch chan int) {
	for i := 0; i < 15; i++ { // 发送完毕后，阻塞
		ch <- i
	}
}

func main() {
	var ok = true
	ch := make(chan int)

	go tel(ch)
	for ok { // 消费者
		i := <-ch // 等待消息，阻塞
		fmt.Printf("ok is %t and the counter is at %d\n", ok, i)
	}
}
