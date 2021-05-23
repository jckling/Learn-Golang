package main

import (
	"fmt"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func main() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName) // 标准输入，空格分隔
	// fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName)

	fmt.Sscanf(input, format, &f, &i, &s)             // 从 input 字符串读取到 f i s
	fmt.Println("From the string we read: ", f, i, s) // From the string we read:  56.12 5212 Go
}
