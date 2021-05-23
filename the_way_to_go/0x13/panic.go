package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting the program")
	panic("A severe error occurred: stopping the program!") // runtime.Error 终止程序
	fmt.Println("Ending the program")
}
