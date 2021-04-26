package main

import "fmt"

func main() {
	sum := 1

	// 省略初始化语句和后置语句时可以省略分号
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
