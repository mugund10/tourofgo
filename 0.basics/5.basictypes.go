package main

import (
	"fmt"
)

// constants 

const(
	tiger = 1 << 2  // leftshift   [  1 which 001 after left shifted two places it is now 00100 which is 4]
   	lion  = 4 >> 2  // rightshift  [ 4 which is 100 after right shift two places it is now 001 which is 1 now]
 )

func main() {

	//basic data type
	var a bool = true
	var b int8 = 127
	var c uint8 = 129
	var d string = "abcde"
	var e float64 = 10.10
	
	//empty assignment

		_ = a
  		_ = b
  		_ = c
  		_ = d
  		_ = e

	//tuple assignment with type coversion

    //b = 127 ,c = 255 after assigning b = -1 , c = 127
	//how b = -1?
	
		b , c = int8(c) , uint8(b)
		fmt.Println(b,c)
	// c= 255, range of b is -128 to 127 totaly 256, from 0 to 127 its same but 
	// after 127 is 128 start as -128 in int8(signed) then 129 -> -127 , so the 
	//last number in uint8 is 255(value of c) and last number in int8 is -1 
	// because 0 to 127 is already mapped.
	// uint8 {0,1,2..127,128,129,...,254,255}
	// int8 {0,1,2,..127,-128,-127,..,-2,-1}
		b , c = int8(c) , uint8(b)
		fmt.Println(b,c)

		
	// numeric constatnts

	fmt.Println(tiger, lion)
	fmt.Printf("the binary forms of tiger is %b which is %[1]d and lion is %b which is %[2]d \n",tiger,lion)
	
	


}
