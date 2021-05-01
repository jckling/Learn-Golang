package main

import "fmt"

func main() {
	w, h := 20, 10	// 打印宽 20 高 10 的矩形
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
