package main

import (
	"container/list"
	"fmt"
)

func main() {
	lst := list.New()
	lst.PushBack(100)
	lst.PushBack(101)
	lst.PushBack(102)
	// fmt.Println("Here is the double linked list:\n", lst)
	for e := lst.Front(); e != nil; e = e.Next() {
		// fmt.Println(e)
		fmt.Println(e.Value)
	}
}
