package main

import (
	"fmt"
	"math"
)

/*
为指针接收者声明方法。
对于某类型 T，接收者的类型可以用 *T 的文法。
T 不能是像 *int 这样的指针。

指针接收者的方法可以修改接收者指向的值。
由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。
*/

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale 方法必须用指针接受者来更改 main 函数中声明的 Vertex 的值
// 若使用值接收者，那么 Scale 方法会对原始 Vertex 值的副本进行操作。
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
