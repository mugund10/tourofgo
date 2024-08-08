package main

import (
	"fmt"
	"time"
)

func main() {

	   method("a")
	go method("ab")
	go method("abc")
	go method("abcd")

	time.Sleep(time.Second)

}

func method(a string) {

	fmt.Println(a)
}
