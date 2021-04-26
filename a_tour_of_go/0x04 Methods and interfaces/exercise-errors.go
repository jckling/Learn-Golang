package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

// 在 Error 方法内调用 fmt.Sprint(e) 会让程序陷入死循环。
// e 是 error 类型，fmt.Sprint(e) 会调用 e.Error()
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("\"cannot Sqrt negative number: %v\"", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := float64(1)
	fmt.Printf("Sqrt approximation of %v:\n", x)
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %v, value = %v\n", i, z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
