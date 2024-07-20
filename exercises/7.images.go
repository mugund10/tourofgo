package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Define the Image type with width and height
type Image struct {
	width, height int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height
}
func (i Image) At(x, y int) color.Color {
	v := uint8((x ^ y) % 256)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	// Create an Image instance with specific dimensions
	m := Image{width: 256, height: 256}
	pic.ShowImage(m)
}
