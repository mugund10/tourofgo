package main

import "fmt"

func main() {
	fmt.Println(sqrt(9))
}

func sqrt(num float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		a := (z*z - num)
		b := (2 * z)
		c := a/b
		z -= c
		fmt.Println(z)
	}
	return z
}
