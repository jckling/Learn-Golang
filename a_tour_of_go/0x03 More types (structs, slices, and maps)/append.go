package main

import "fmt"

/*
内建函数 append 可以为切片追加新的元素；
append 的结果是一个包含原切片所有元素加上新添加元素的切片。

func append(s []T, vs ...T) []T
s 是一个元素类型为 T 的切片
其余类型为 T 的值将会追加到该切片的末尾

当 s 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。
返回的切片会指向这个新分配的数组。
*/

func main() {
	var s []int
	printSlice(s)

	// 添加一个空切片
	s = append(s, 0)
	printSlice(s)

	// 这个切片会按需增长
	s = append(s, 1)
	printSlice(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
