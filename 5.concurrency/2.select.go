package main

import (
	"fmt"
	"time"
)

func main() {
	tt := time.Now()
	c := make(chan int)
	quit := make(chan int)

	defer func(){
		fmt.Println(time.Since(tt))
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) //receives val x
		}
		quit <- 0 //sends 0
	}()
	fib(c, quit)


	//default in select

	tick := time.Tick(time.Second*2)
	boom := time.After(time.Second*3)
	for {
		select {
		case <-tick : fmt.Println("tick")
		case <-boom : fmt.Println("boom")
		default : fmt.Println(" .")
		time.Sleep(1 * time.Second)
		}
	}

}

func fib(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: //sends val x
			x, y = y, x+y
		case <-quit : //receives 0
			fmt.Println("quits")
			return
		}
	}
}
