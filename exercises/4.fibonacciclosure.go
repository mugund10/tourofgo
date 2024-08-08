package main

import (
	"fmt"
	"time"
)

func main()  {
	tt := time.Now()
	defer func(){
		fmt.Println(time.Since(tt))
	}()
	newfib := fibo()
	for i := 0 ; i <10 ; i++ {
		fmt.Println(newfib())
	}
}

func fibo()func()int{
		a , b  := 1, 1
	return func() int{
		result:= a
		a ,b = b , a+b	
        return result
	}
}