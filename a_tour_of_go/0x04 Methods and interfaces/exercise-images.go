package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

/*
type Image interface {
	// ColorModel returns the Image's color model.
	ColorModel() color.Model
	// Bounds returns the domain for which At can return non-zero color.
	// The bounds do not necessarily contain the point (0, 0).
	Bounds() Rectangle
	// At returns the color of the pixel at (x, y).
	// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
	// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
	At(x, y int) color.Color
}
*/

type Image struct {
	width  int
	height int
}

// ColorModel 应当返回图像的颜色模型。
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds 应当返回一个 image.Rectangle 。
// 例如 image.Rect(0, 0, w, h)。
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

// At 应当返回一个颜色。
// 上一个图片生成器的值 v 对应于此次的 color.RGBA{v, v, 255, 255}。
func (img Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
