package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	fName := "D:/Study/Github/Learn-Golang/the_way_to_go/0x12/MyFile.gz"
	var r *bufio.Reader

	fi, err := os.Open(fName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], fName,
			err)
		os.Exit(1)
	}
	defer fi.Close()

	fz, err := gzip.NewReader(fi)
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}
