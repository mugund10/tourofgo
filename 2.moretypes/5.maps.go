package main

import "fmt"

func main() {
	name1 := make(map[int]string)
	fmt.Println(name1)

	name1[1] = "ONE"
	name1[2] = "TWO"
	fmt.Println(name1[1])
	fmt.Println("the value of name1[2] after delte function",name1[2],)
	delete(name1,2)
	fmt.Println("the value of name1[2] after delte function",name1[2],)
	
	newname1, ok := name1[1]
	if !ok {
		fmt.Println("no values")
	}else {
		fmt.Println(newname1)
	}

	newname2, ok := name1[2]
	if !ok {
		fmt.Println("no values")
	}else {
		fmt.Println(newname2)
	}
	

}
