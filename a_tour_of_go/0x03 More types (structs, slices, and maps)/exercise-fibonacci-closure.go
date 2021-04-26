package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	f1, f2 := 0, 1
	return func() int {
		f1, f2 = f2, f1+f2
		return f1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
