package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	pixel := make([][]uint8, dy)
	for i := range pixel {
		pixel[i] = make([]uint8, dx)
		for j := range pixel {
			pixel[i][j] = uint8((i + j) / 2)
			//pixel[i][j] = uint8(i * j)
			//pixel[i][j] = uint8(i % (j + 1))
		}
	}
	return pixel
}

func main() {
	pic.Show(Pic)
}
