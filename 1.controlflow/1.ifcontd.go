package main

import "fmt"

func main() {

	// here a is shortstatement, which means
	// its declared inside the if condition 
	// and cant be accessed by outside the if statement
	if a := 10; a < 20 {
		fmt.Println(a)

	} else if a == 22 {
		fmt.Println	(a)	
	} 
	//fmt.Println(a) its gives error because a 
	// initialized inside the if statement
}
