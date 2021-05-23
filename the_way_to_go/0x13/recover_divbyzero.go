package main

import (
	"fmt"
)

func badCall() {
	a, b := 10, 0
	n := a / b     // 除零异常，引发 panic
	fmt.Println(n) // 不会执行到这里
}

func test() {
	defer func() {
		if e := recover(); e != nil { // 处理错误
			fmt.Printf("Panicing %s\r\n", e)
		}

	}()
	badCall()
	fmt.Printf("After bad call\r\n") // 不会执行到这里
}

func main() {
	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")
}
