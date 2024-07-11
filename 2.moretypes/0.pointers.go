package main

import "fmt"

func main() {

	i, j := 10, 100

	p := &i

	fmt.Println("where the both statement are contains a address of a value",p , &i)

	*p = j 

	fmt.Println("where *p can change the value of i because * denotes underlying value",j,*p,i)

	a := 10 
	b := &a
	c := &b
	d := &c

	//now to change a value of 'a' by only using 'd'

	***d = 12
//  **(*d)= c | *(*c)= b | *b = a
	fmt.Println(a)
}
