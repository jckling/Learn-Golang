package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	CopyFile("D:/Study/Github/Learn-Golang/the_way_to_go/0x12/target.txt", "D:/Study/Github/Learn-Golang/the_way_to_go/0x12/source.txt")
	fmt.Println("Copy done!")
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644) // 文件，模式，权限
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
