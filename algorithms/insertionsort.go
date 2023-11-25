package sort

func InsertionSort(arr []int) {
	N := len(arr)
	for i := 1; i < N; i++ {
		for j := i; j > 0 && less(arr[j], arr[j-1]); j-- {
			exchange(arr, j, j-1)
		}
	}
}
