package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runing on ")
	//initializing variable on switch will be accesed by switch only
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("DARWIN")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s \n", os)
		// inside	fmt.Println(os)
	}
	// outside switch fmt.Println(os)
	i := 12
	switch i {
	case neew(i):
		fmt.Println("11")
	case 12:
		fmt.Println("12")
	//it doesnt go throught all possibilities
	//other than it will go the for one which is on top
	case neew(i + 1):
		fmt.Println("12 on func")
		// thats why the upper case not gets executed
	}
	// switch without a condittion will acts as
	// switch true
	switch {

	case i > 20:
		fmt.Println("i is greater than 20")
	case i == 20:
		fmt.Println("i is equal to 20")
	case i < 20:
		fmt.Println("i is less than 20")

	}

}

func neew(a int) int {
	return a - 1
}
