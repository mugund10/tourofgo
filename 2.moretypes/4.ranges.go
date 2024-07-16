package main

import "fmt"

func main() {
	sli := []int{1, 2, 3, 4, 5, 6}
	for key, value := range sli {
		fmt.Println("key :", key, " value :", value)
	}
}
