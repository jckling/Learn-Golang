package main

import (
	"fmt"
)

// 近似平方根
func Sqrt(x float64) float64 {
	// z := 1.0
	z := float64(1)
	fmt.Printf("Sqrt approximation of %v:\n", x)
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z) // 牛顿法
		fmt.Printf("Iteration %v, value = %v\n", i, z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
