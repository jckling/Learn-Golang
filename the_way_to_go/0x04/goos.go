package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	var goos string = runtime.GOOS	// 运行时获取操作系统类型
	fmt.Printf("The operating system is: %s\n", goos)	// 格式化输出
	path := os.Getenv("PATH")		// 获取环境变量的值
	fmt.Printf("Path is %s\n", path)
}
