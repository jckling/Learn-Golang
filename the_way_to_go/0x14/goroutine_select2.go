package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// setting GOMAXPROCS to 2 gives +- 22% performance increase,
	// but increasing the number doesn't increase the performance
	// without GOMAXPROCS: +- 86000
	// setting GOMAXPROCS to 2: +- 105000
	// setting GOMAXPROCS to 3: +- 94000
	runtime.GOMAXPROCS(2)
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)

	time.Sleep(1e9)
}

// 生产者
func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

// 生产者
func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

// 消费者
func suck(ch1, ch2 chan int) {
	for i := 0; ; i++ { // 计数
		select {
		case v := <-ch1:
			fmt.Printf("%d - Received on channel 1: %d\n", i, v)
		case v := <-ch2:
			fmt.Printf("%d - Received on channel 2: %d\n", i, v)
		}
	}
}
