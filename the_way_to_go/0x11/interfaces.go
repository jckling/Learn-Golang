package main

import "fmt"

// 接口
type Shaper interface {
	Area() float32		// 方法
	// Perimeter() float32
}


// 结构体
type Square struct {
	side float32
}

// 接受者类型 Square
// 结构体 Square 实现了接口 Shaper
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	sq1 := new(Square)
	sq1.side = 5

	var areaIntf Shaper
	areaIntf = sq1		// 赋值给接口类型的变量
	// shorter, without separate declaration:
	// areaIntf := Shaper(sq1)
	// or even:
	// areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
}
