package main

import "fmt"

/*
切片并不存储任何数据，它只是描述了底层数组中的一段。
更改切片的元素会修改其底层数组中对应的元素。
与它共享底层数组的切片都会观测到这些修改。
*/

func main() {
	// 数组
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	// 切片
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// 修改数据，所有的都被修改
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
