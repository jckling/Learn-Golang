package main

import "fmt"

/*
函数可以没有参数或接受多个参数。

注意类型在变量名之后。
*/

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
