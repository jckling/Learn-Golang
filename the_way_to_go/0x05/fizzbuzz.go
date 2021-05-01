package main

import "fmt"

const (
	FIZZ     = 3	// 3 的倍数
	BUZZ     = 5	// 5 的倍数
	FIZZBUZZ = 15	// 3 和 5 的倍数
)

func main() {
	for i := 0; i <= 100; i++ {
		switch {
		case i%FIZZBUZZ == 0:
			fmt.Println("FizzBuzz")
		case i%FIZZ == 0:
			fmt.Println("Fizz")
		case i%BUZZ == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
