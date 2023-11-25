package search

// sorted array
func BinarySearch(key int, arr []int) int {
	lo := 0
	hi := len(arr) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if key < arr[mid] {
			hi = mid - 1
		} else if key > arr[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
