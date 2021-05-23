package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// go run cat_numbered.go -n products.txt
var numberFlag = flag.Bool("n", false, "number each line") // -n 标志位

func cat(r *bufio.Reader) {
	i := 1
	for {
		buf, err := r.ReadBytes('\n')

		if err == io.EOF {
			break
		}

		if *numberFlag { // 打印行号
			fmt.Fprintf(os.Stdout, "%5d %s", i, buf)
			i++
		} else {
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
	}
	return
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 { // go run cat_numbered.go
		cat(bufio.NewReader(os.Stdin)) // 标准输入
	}

	for i := 0; i < flag.NArg(); i++ { // go run cat_numbered.go products.txt
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
		f.Close()
	}
}

// go run cat_numbered.go -n
