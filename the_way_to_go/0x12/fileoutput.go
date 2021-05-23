package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var outputWriter *bufio.Writer
	// var outputFile *os.File
	// var outputError os.Error
	// var outputString string

	// 文件名，标志，权限
	outputFile, outputError := os.OpenFile("D:/Study/Github/Learn-Golang/the_way_to_go/0x12/output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"

	for i := 0; i < 10; i++ { // 写入十次
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}
