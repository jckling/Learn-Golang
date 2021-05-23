package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "Alice "
	if len(os.Args) > 1 { // 命令行参数
		who += strings.Join(os.Args[1:], " ") // os.Args[0] 程序名
	}
	fmt.Println("Good Morning", who)
}

/*
go run os_args.go
output: Good Morning Alice

go run os_args.go jck
output: Good Morning Alice jck
*/
