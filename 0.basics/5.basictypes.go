package main

import (
	"fmt"
)

type kings []byte

func main() {
	var a bool
	fmt.Println(a)
	var sign int8 = -128
	var unsign uint8 = 255
	var total int16 = int16(unsign) + int16(sign)
	fmt.Println(total)
	var newguy kings
	newguy = []byte{byte(total)+byte(total)}
	fmt.Println(string(newguy))
	


}
