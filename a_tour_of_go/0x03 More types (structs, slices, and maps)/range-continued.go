package main

import "fmt"

/*
可以将下标或值赋予 _ 来忽略它。
for i, _ := range pow
for _, value := range pow

若只需要索引，忽略第二个变量即可。
for i := range pow
*/

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
