package main

import "fmt"


func main() {
	defer fmt.Println("world")
// defers will run the function after all functions gets executed 
	fmt.Println("hello")

	// actually defer is like stack
	//so its in the order of lifo
	//"last in first out"
	
	for key , val := range "hello" {
		defer fmt.Println(key, string(val))
	}

}