package main

import (
	"fmt"
	"math"
)

// 在 if 的简短语句中声明的变量同样可以在任何对应的 else 块中使用。

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// 也可以在 else 块中使用 v
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
