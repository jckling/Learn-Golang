package sort

func ShellSort(arr []int) {
	N := len(arr)
	h := 1
	for ; h < N/3; h = 3*h + 1 {
	}

	for ; h >= 1; {
		for i := h; i < N; i++ {
			for j := i; j >= h && less(arr[j], arr[j-h]); j -= h {
				exchange(arr, j, j-h)
			}
		}
		h = h / 3
	}
}
