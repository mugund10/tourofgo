package main

import "fmt"

func main() {

	for i , j :=  3, 9  ; i ^ j > 9  ; i++{
          fmt.Println(i,j,i ^ j )
	}

a := 10

b := 12

fmt.Println(a &^ b) 

}