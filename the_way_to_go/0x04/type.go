package main
import "fmt"

type TZ int
type Rope string

func main() {
	var a, b TZ = 3, 4
	c := a + b
	fmt.Printf("c has the value: %d\n", c)

	var s Rope = "newline"
	fmt.Printf("%s", s)
}