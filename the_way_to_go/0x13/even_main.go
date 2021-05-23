package main

import (
	"../0x09/even"
	"fmt"
)

func main() {
	for i := 0; i <= 100; i++ { // 0~100 判断奇数
		fmt.Printf("Is the integer %d even? %v\n", i, even.Even(i))
	}
}
