package main

import (
	"fmt"
	"os"
)

func tel(ch chan int, quit chan bool) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	quit <- true
}

func main() {
	var ok = true
	ch := make(chan int)    // 传递数据
	quit := make(chan bool) // 是否关闭

	go tel(ch, quit)
	for ok {
		select {
		case i := <-ch:
			fmt.Printf("The counter is at %d\n", i)
		case <-quit:
			os.Exit(0)
		}
	}
}
