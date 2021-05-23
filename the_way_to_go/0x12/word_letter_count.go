package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var nrchars, nrwords, nrlines int

func main() {
	nrchars, nrwords, nrlines = 0, 0, 0
	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter some input, type S to stop: ")

	for {
		input, err := inputReader.ReadString('\n') // 英文字母容易统计
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
		}

		if input == "S\n" { // 单行 S+回车 终止
			fmt.Println("Here are the counts:")
			fmt.Printf("Number of characters: %d\n", nrchars)
			fmt.Printf("Number of words: %d\n", nrwords)
			fmt.Printf("Number of lines: %d\n", nrlines)
			os.Exit(0)
		}

		Counters(input)
	}
}

func Counters(input string) {
	nrchars += len(input) - 2 // -2 for \r\n
	nrwords += len(strings.Fields(input))
	nrlines++
}
