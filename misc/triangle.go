package main

import "fmt"


func main() {
 tri(6)

}

func tri(val int){
	for i:=1 ; i <= val; i++ {
		fmt.Print("*")
		for j:=1 ; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}