package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := ""
	if len(os.Args) > 1 { // go run hello_who.go jck
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Printf("Hello %s!\n", who)
}
