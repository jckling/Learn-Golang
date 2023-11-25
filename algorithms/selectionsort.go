package sort

import "fmt"

func less(a int, b int) bool {
	return a < b
}

func exchange(arr []int, i int, min int) {
	if i == min {
		return
	}
	a := arr[i]
	arr[i] = arr[min]
	arr[min] = a
}

func Show(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func SelectionSort(arr []int) {
	N := len(arr)
	for i := 0; i < N; i++ {
		min := i
		for j := i + 1; j < N; j++ {
			if less(arr[j], arr[min]) {
				min = j
			}
		}
		exchange(arr, i, min)
	}
}
