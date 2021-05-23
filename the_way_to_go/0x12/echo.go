package main

import (
	"flag" // command line option parser
	"os"
)

var NewLine = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.PrintDefaults() // 打印标志位信息
	flag.Parse()         // Scans the arg list and sets up flags
	var s string = ""
	for i := 0; i < flag.NArg(); i++ { // 返回参数的数量，不包含程序名，从 0 开始
		if i > 0 {
			s += " "
			if *NewLine { // -n is parsed, flag becomes true
				s += Newline // 添加换行符
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}

/*
	// case 1: go run echo.go A B C
	  -n    print newline
	A B C

	// case 2: go run echo.go -n A B C
	  -n    print newline
	A
	B
	C
*/
