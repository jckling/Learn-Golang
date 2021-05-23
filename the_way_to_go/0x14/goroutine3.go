package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	close(ch) // 关闭通道
}

func getData(ch chan string) {
	for {
		input, open := <-ch // 阻塞，等待通道消息
		if !open {	// 检测通道是否关闭
			break
		}
		fmt.Printf("%s ", input)
	}
}
