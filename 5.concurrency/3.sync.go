package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main(){
// wait group

	
	wg.Add(2)
	go newe()
	go newe()

	wg.Wait()
	fmt.Println("main go routine")
	
// mutex 

iterations := 1000
sum := 0

wg.Add(iterations)
var mu sync.Mutex

for i := 0 ; i < iterations ; i++ {
	go func(){
		mu.Lock()
		sum++
		mu.Unlock()
		wg.Done()
	}()

}
wg.Wait()
fmt.Println(sum)

}


func newe(){
	fmt.Println("new1 go routine")
		wg.Done()
	}