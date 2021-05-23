package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

var content string
var vc VCard

func main() {
	// using a decoder:
	file, _ := os.Open("D:/Study/Github/Learn-Golang/the_way_to_go/0x12/vcard.gob")
	defer file.Close()

	inReader := bufio.NewReader(file) // 读取文件
	dec := gob.NewDecoder(inReader)   // 解析 go binary
	err := dec.Decode(&vc)
	if err != nil {
		log.Println("Error in decoding gob")
	}
	fmt.Println(vc)
}
