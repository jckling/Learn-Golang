package main

import (
	"fmt"
	"time"
)

func main() {
	suck(pump())
	time.Sleep(1e9)
}

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	go func() {
		for v := range ch { // 从通道中获取值
			fmt.Println(v)
		}
	}()
}
