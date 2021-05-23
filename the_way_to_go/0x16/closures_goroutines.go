package main

import (
	"fmt"
	"time"
)

var values = [5]int{10, 11, 12, 13, 14}

func main() {
	// version A:
	for ix := range values { // ix 是索引
		func() {
			fmt.Print(ix, " ")
		}() // call closure, prints each index
	}
	fmt.Println()

	// version B: same as A, but call closure as a goroutine
	for ix := range values { // ix 是索引
		go func() {
			fmt.Print(ix, " ") // 每次循环都打印最后一个值，因此协程可能在循环结束后还没有开始执行
		}()
	}
	time.Sleep(1e9)

	// version C: the right way
	for ix := range values { // ix 是索引
		go func(ix interface{}) {
			fmt.Print(ix, " ") // ix 每次循环都被重新赋值，输出取决于协程何时被执行
		}(ix)
	}
	time.Sleep(1e9)

	// version D: print out the values:
	for ix := range values {
		val := values[ix] // 循环体内部声明
		go func() {
			fmt.Print(val, " ")
		}()
	}
	time.Sleep(1e9)
}

/* Output:
0 1 2 3 4
4 4 4 4 4
0 2 4 1 3
14 10 11 12 13
*/
