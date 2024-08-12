package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	defer func() {
		fmt.Println("time taken : ", time.Since(start))
	}()
	//channels
	newc := make(chan string)
	go call("a", newc)
	go call("b", newc)
	go call("c", newc)
	go call("d", newc)
	fmt.Println(<-newc)
	fmt.Println(<-newc)
	fmt.Println(<-newc)
	fmt.Println(<-newc)

	//channels with properties
	sender := (chan<- string)(newc)
	receiver := (<-chan string)(newc)
	go tcall("abcd", sender)
	fmt.Println("receiver",<-receiver)

	//range or channel iteration
	namelist := []string{"abc", "def", "ghi", "jkl"}
	namec := make(chan string)
	go rangee(namelist, namec)
	// for val := range namec {
	// 	fmt.Println(val)
	// }
	// or
	for i := 0; i < len(namelist); i++ {
		fmt.Println(<-namec)
	}

	//buffered channels
	bc := make(chan string, 4)
	bc <- "a"
	bc <- "b"
	bc <- "c"
	//bc <- "d"
	fmt.Println(<-bc,<-bc,<-bc)


}

func call(a string, cha chan string) {

	fmt.Println("call called")
	cha <- a

}

func tcall(a string, cha chan<- string) {
	cha <- "tcall"
}

func rangee(a []string, cha chan string) {
	for _, as := range a {
		cha <- as
	}
	close(cha)
}
