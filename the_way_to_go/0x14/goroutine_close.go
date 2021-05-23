package main

import (
	"fmt"
)

func tel(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	close(ch) // 关闭通道
}

func main() {
	var ok = true
	var i int
	ch := make(chan int)

	go tel(ch)
	for ok {
		if i, ok = <-ch; ok { // 判断通道是否关闭
			fmt.Printf("ok is %t and the counter is at %d\n", ok, i)
		}
	}
}
