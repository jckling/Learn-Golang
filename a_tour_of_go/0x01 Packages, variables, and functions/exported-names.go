package main

import (
	"fmt"
	"math"
)

/*
在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的。

在导入一个包时，只能引用其中已导出的名字。
任何“未导出”的名字在该包外均无法访问。
*/

func main() {
	//fmt.Println(math.pi) // 编译错误！
	fmt.Println(math.Pi)
}
