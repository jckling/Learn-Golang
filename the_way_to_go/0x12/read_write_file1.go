package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	inputFile := "D:/Study/Github/Learn-Golang/the_way_to_go/0x12/products.txt"
	outputFile := "D:/Study/Github/Learn-Golang/the_way_to_go/0x12/products_copy.txt"
	buf, err := ioutil.ReadFile(inputFile) // 读取整个文件的内容
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644) // oct, not hex 写入文件
	if err != nil {
		panic(err.Error())
	}
}
