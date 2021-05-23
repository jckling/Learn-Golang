package main

import (
	"fmt"
)

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil { // 4
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)                                     // 引发 panic
	fmt.Println("Returned normally from g.") // 不会执行到这里
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i)) // 4
	}
	defer fmt.Println("Defer in g", i) // 返回时运行
	fmt.Println("Printing in g", i)
	g(i + 1) // 递归调用
}
