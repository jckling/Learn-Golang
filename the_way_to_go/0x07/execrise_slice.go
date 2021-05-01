package main

import "fmt"

func main() {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(b[1:4])
	fmt.Println(b[:2])
	fmt.Println(b[2:])
	fmt.Println(b[:])

	s := make([]byte, 5)
	fmt.Println(len(s), cap(s))
	s = s[2:4]
	fmt.Println(len(s), cap(s))

	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]
	fmt.Println(s2)
	s2[1] = 't'
	fmt.Println(s1, s2)

	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var a = ar[5:7] // reference to subarray {5,6} - len(a) is 2 and cap(a) is 5
	a = a[0:4]      // ref of subarray {5,6,7,8} - len(a) is now 4 but cap(a) is still 5
	fmt.Println(len(a[4:4]))
	fmt.Println(len(a[4:5]))
}
