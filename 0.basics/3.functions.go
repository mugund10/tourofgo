package main

import "fmt"

func main() {
	fmt.Println(newAdd(newOne(),newOne()))  //(1+1) = 2
	fmt.Println(newSwap("FIRST","LAST"))   //LAST FIRST
	fmt.Println(newNamed(10))             //7 3
}

func newOne() int {
	return 1
}

func newAdd(a, b int) int {
	return a + b
}

func newSwap(a,b string) (string,string) {
	return b, a
}

func newNamed(a int) (x,y int) { //we name x and y as the returning variable of type int
x = a-3
y = a-x
return              // its the benefit of named return.
					
}