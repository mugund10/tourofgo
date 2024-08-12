package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	waitgroup()
	muutex()
	oonce()
}

func waitgroup() {
	//waitgroup

	wg.Add(3) // how many time we have to wait

	go rout()

	wg.Wait()
	fmt.Println("waitgroup func start")

	

}

func rout() {
	go rout1()
	go rout2()
	fmt.Println("rout function")
	wg.Done()

}

func rout1() {
	fmt.Println("rout1 function")
	wg.Done()
}

func rout2() {
	fmt.Println("rout2 function")
	wg.Done()

	wg.Add(2)
	go rout1()
	go rout1()
	wg.Wait()
	fmt.Println("rout2 function end")
}


var sum int
var mu sync.Mutex

//helps with mutual exclusion of resources among different go routines
func muutex(){

	iteration := 1000

	wg.Add(iteration)
	for i := 0 ; i < iteration; i++{
		
		go inc()
		
	}
	wg.Wait()
	fmt.Println(sum)


}

func inc(){
	mu.Lock()
	sum++
	mu.Unlock()
	wg.Done()
}



var sum2 int
var on sync.Once

func oonce(){



	
	wg.Add(1)
	go onlyone()
	wg.Wait()

	fmt.Println(sum2)
	

}


func onlyone(){

	on.Do(increonce)
	wg.Done()

}

func increonce(){
	sum2++
	
}