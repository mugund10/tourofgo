package main

import "fmt"

func main() {
	fmt.Println(sqrt(2))
}

func sqrt(num float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		a := (z*z - num)
		fmt.Println("a = ",a)
		b := (2 * z)
		fmt.Println("b = ",b)
		c := a/b
		fmt.Println("c = ",c)
		z -= c
		fmt.Println("z = ",z)
	}
	return z
}
