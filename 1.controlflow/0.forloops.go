package main

import (
	"fmt"
)

func main() {

	for i, j := 3, 9; i^j > 9; i++ {
		fmt.Println(i, j, i^j)
	}

	a := 3

	b := 12

	fmt.Println(a &^ b)

	// go's while loop

	for a < b {
		a += a
	}

	fmt.Println(b)

// for {}

}
