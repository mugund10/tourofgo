package main

import "fmt"

type tipes interface{
	int | float64 | string
}

func main(){
	fmt.Println(adder(10,20))
	fmt.Println(adder("a","b"))
}

func adder[typ tipes](a... typ) typ {
	var num typ
	for val := range a {
		num += a[val]
	}
	return num
}