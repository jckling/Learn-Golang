package main

import "fmt"

/*
切片的零值是 nil 。
nil 切片的长度和容量为 0 且没有底层数组。
*/

func main() {
	var s []int

	if s == nil {
		fmt.Println("nil!")
	}

	fmt.Println(s, len(s), cap(s))
}
