package main

import "fmt"

/*
nil 接口值既不保存值也不保存具体类型。
为 nil 接口调用方法会产生运行时错误，因为接口的元组内并未包含能够指明该调用哪个具体方法的类型。
*/

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M() // 编译错误！
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
