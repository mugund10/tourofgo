package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	sli := make([][]uint8, dy)

	for keyX, _ := range sli {
		sli1 := make([]uint8, dx)
		for keyY, _ := range sli1 {
			sli1[keyY] = uint8((keyX + keyY) / 2)
		}
		sli[keyX] = sli1
	}

	return sli

}

func main() {
	pic.Show(Pic)

}
