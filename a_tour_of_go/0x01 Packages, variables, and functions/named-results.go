package main

import "fmt"

/*
Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。
返回值的名称应当具有一定的意义，它可以作为文档使用。
没有参数的 return 语句返回已命名的返回值，也就是直接返回。

直接返回语句应当仅用在短函数中。
在长的函数中它们会影响代码的可读性。
*/

// 返回已命名的返回值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
