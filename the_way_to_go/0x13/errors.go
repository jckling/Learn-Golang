package main

import (
	"errors"
	"fmt"
)

var errNotFound error = errors.New("Not found error") // 定义错误类型

func main() {
	fmt.Printf("error: %v", errNotFound)
}
