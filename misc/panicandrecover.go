package main

import "fmt"

func main() {
	const k = 0
	defer first()
	if k != 1 {
		fmt.Println("this will 1")
		panic("impossible")
		fmt.Println("this wont")
	}
	fmt.Println("this wont 2")
}

func first() {

	if r := recover(); r != nil {
		fmt.Println("recovered")
	}
}
