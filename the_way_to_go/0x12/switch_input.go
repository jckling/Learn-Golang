package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	fmt.Printf("Your name is %s", input)

	// version 1:
	switch input {
	case "Philip\n":
		fmt.Println("Welcome Philip!")
	case "Ivo\n":
		fmt.Println("Welcome Ivo!")
	case "Chris\n":
		fmt.Println("Welcome Chris!")
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}

	// version 2:
	switch input {
	case "Philip\n":
		fallthrough
	case "Ivo\n":
		fallthrough
	case "Chris\n":
		fmt.Printf("Welcome %s", input) // 输入本来就有换行符
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}

	// version 3:
	switch input {
	case "Philip\n", "Ivo\n":
		fmt.Printf("Welcome %s", input) // 输入本来就有换行符
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}
}

/*
	// case 1: Philip
	Your name is Philip
	Welcome Philip!
	Welcome Philip
	Welcome Philip

	// case 2: Ivo
	Your name is Ivo
	Welcome Ivo!
	Welcome Ivo
	Welcome Ivo

	// case 3: Chris
	Your name is Chris
	Welcome Chris!
	Welcome Chris
	You are not welcome here! Goodbye!

	// case 4: Jck
	Your name is Jck
	You are not welcome here! Goodbye!
	You are not welcome here! Goodbye!
	You are not welcome here! Goodbye!
*/
