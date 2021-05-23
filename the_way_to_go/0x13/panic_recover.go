package main

import (
	"fmt"
)

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if e := recover(); e != nil { // bad end
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	badCall()                        // 引发 panic
	fmt.Printf("After bad call\r\n") // 不会执行到这里
}

func main() {
	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")
}
