package main

import (
	"fmt"
	"strings"
)

func main() {
	asciiOnly := func(c rune) rune {
		if c > 127 {	// 非 ASCII 字符替换成空格
			return ' '
		}
		return c
	}
	fmt.Println(strings.Map(asciiOnly, "Jérôme Österreich"))
}
