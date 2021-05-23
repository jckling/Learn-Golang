package even

func Even(i int) bool { // exported function
	return i%2 == 0
}

func Odd(i int) bool { // exported function
	return i%2 != 0
}
