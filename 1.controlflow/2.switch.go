package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runing on ")
	switch os := runtime.GOOS; os {
	case "darwin" : fmt.Println("DARWIN")
	case "linux" : fmt.Println("linux")
	default : fmt.Printf("%s \n",os)
    //switch inside switch
	i := 12
	switch i {
	case neew(i) : fmt.Println("11")
	case 12 : fmt.Println("12")
	case neew(i+1) : fmt.Println("12 on func")
	}
}

}

func neew(a int)int{
return a - 1
}