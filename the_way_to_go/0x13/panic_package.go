package main

import (
	"./parse"
	"fmt"
)

func main() {
	var examples = []string{
		"1 2 3 4 5",
		"100 50 25 12.5 6.25",
		"2 + 2 = 4",
		"1st class",
		"",
	}

	for _, ex := range examples {
		fmt.Printf("Parsing %q:\n  ", ex)
		nums, err := parse.Parse(ex) // 解析整数
		if err != nil {
			fmt.Println(err) // here String() method from ParseError is used
			continue
		}
		fmt.Println(nums)
	}
}
