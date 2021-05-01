package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v) // 始终为初始化的默认值 0
		v = 5
	}

	// 无限循环，随机值
	//for i := 0; ; i++ {
	//	fmt.Println("Value of i is now:", i)
	//}

	// 无限循环，0
	//for i := 0; i < 3; {
	//	fmt.Println("Value of i:", i)
	//}

	s := ""
	for ; s != "aaaaa"; {
		fmt.Println("Value of s:", s) // 空值,a,aa,aaa,aaaa
		s = s + "a"
	}

	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
		// 0,5,a
		// 1,6,aa
		// 2,7,aaa
	}

	i := 0
	for { //since there are no checks, this is an infinite loop
		if i >= 3 {
			break
		}
		//break out of this for loop when this condition is met
		fmt.Println("Value of i is:", i)
		i++
	}
	fmt.Println("A statement just after for loop.")

	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd:", i)
	}
}
