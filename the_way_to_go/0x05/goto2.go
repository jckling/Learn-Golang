package main

import "fmt"

func main() {
	a := 1
	goto TARGET // compile error
	b := 9
TARGET:
	b += a
	fmt.Printf("a is %v *** b is %v", a, b)
}
