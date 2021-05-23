package main

import (
	"../0x11/stack"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 逆波兰式
func main() {
	buf := bufio.NewReader(os.Stdin)
	calc1 := new(stack.Stack)
	fmt.Println("Give a number, an operator (+, -, *, /), or q to stop:")

	for {
		token, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println("Input error !")
			os.Exit(1)
		}

		token = token[0 : len(token)-1] // remove "\n"
		// fmt.Printf("--%s--\n",token)  // debug statement

		switch {
		case token == "q": // stop als invoer = "q"
			fmt.Println("Calculator stopped")
			return

		case token >= "0" && token <= "999999":
			i, _ := strconv.Atoi(token)
			calc1.Push(i)

		case token == "+":
			q, _ := calc1.Pop()
			p, _ := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)+q.(int))

		case token == "-":
			q, _ := calc1.Pop()
			p, _ := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)-q.(int))

		case token == "*":
			q, _ := calc1.Pop()
			p, _ := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)*q.(int))

		case token == "/":
			q, _ := calc1.Pop()
			p, _ := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)/q.(int))

		default:
			fmt.Println("No valid input")
		}
	}
}

/*
	输入：
	3
	4
	+

	输出：
	The result of 3 + 4 = 7
*/
