package main

import (
	s "./sort"
	"fmt"
)

func main() {
	array := []int{18, 20, 15, 22, 16}
	s.SelectionSort(array)
	s.Show(array)

	fmt.Println("====")
	array2 := []int{1, 2, 5, 22, 16}
	s.InsertionSort(array2)
	s.Show(array2)

	fmt.Println("====")
	array3 := []int{11, 32, 5, 22, 16}
	s.ShellSort(array3)
	s.Show(array3)
}
