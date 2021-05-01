package main

import (
	"./float64"
	"fmt"
)

func main() {
	f1 := float64.NewFloat64Array()
	f1.Fill(10)
	fmt.Printf("Before sorting %s\n", f1)
	float64.Sort(f1)
	fmt.Printf("After sorting %s\n", f1)
	if float64.IsSorted(f1) {
		fmt.Println("The float64 array is sorted!")
	} else {
		fmt.Println("The float64 array is NOT sorted!")
	}
}
