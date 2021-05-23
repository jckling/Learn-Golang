package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var input string
var err error

func main() {
	// 创建读取器，与标准输入绑定
	inputReader = bufio.NewReader(os.Stdin) // reader for input
	fmt.Println("Please enter some input: ")
	input, err = inputReader.ReadString('\n') // 读取到 \n 或文件结尾 io.EOF

	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}
