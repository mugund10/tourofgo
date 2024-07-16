package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	//like hypot(3,4)
	return fn(3, 4)
}
func main() {
	//like func hypot(a,b float64) float64 {return math.Sqrt(a*a + b*b)}
	hypot := func(a, b float64) float64 {
		return math.Sqrt(a*a + b*b)
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow)) //math.Pow(3,4)

	// funcs adds two numbers
	newpot1 := func(a, b int) int {
		return a + b
	}
	// funcs multiplies two numbers
	newpot2 := func(a, b int) int {

		return a * b
	}

	fmt.Println(hukkum(newpot1))
	fmt.Println(hukkum(newpot2))

	//closures
	newclosu := clos()

	for i := 0; i < 9; i++ {
		fmt.Println("closures",newclosu(i))
	}

}

func hukkum(abd func(int, int) int) int {
	return abd(3, 9)
}

// closure example
func clos() func(int) int {
	sum := 0
	//the return function accessing the variable sum
	return func(x int) int {
		sum += x
		return sum
	}

}
