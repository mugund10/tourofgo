package main

import (
	"fmt"
)

type ip struct {
	addr [4]byte
}

type school struct {
	name    string
	address string
}

func (i ip) String() string {
	a := &i.addr
	return fmt.Sprintf("%v.%v.%v.%v",a[0],a[1],a[2],a[3] )
}

func (s school) String() string {
	return fmt.Sprintf("the %v school is located in the %v",s.name,s.address)
}

func main() {
	google := ip{[4]byte{8, 8, 8, 8}}
	fmt.Println(google)
	microsoft := school{"micros","us"}
	fmt.Println(microsoft)
}
