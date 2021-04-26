package main

import "fmt"

/*
即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。
在 Go 中通常会写一些方法来优雅地处理它（如本例中的 M 方法）。
注意: 保存了 nil 具体值的接口其自身并不为 nil。
*/

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
