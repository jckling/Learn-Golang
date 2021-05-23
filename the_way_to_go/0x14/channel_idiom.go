package main

import (
	"fmt"
	"time"
)

func main() {
	stream := pump()
	go suck(stream)
	time.Sleep(1e9)
}

func pump() chan int {
	ch := make(chan int)
	go func() { // 匿名函数
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch // 返回通道（工厂）
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
