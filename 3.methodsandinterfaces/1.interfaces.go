package main

import "fmt"

type mugu string

type bala struct {
	x, y string
}
// from interface Stringer
func (b bala) String() string {
	return fmt.Sprintf("the %v and %v", b.x, b.y)
}

type Mb interface {
	Prints()
}

func (m mugu) Prints() {
	fmt.Println(m)
}

func (b bala) Prints() {
	fmt.Println(b)
}

func main() {
	var mb Mb

	m := mugu("mugund10")
	b := bala{"ba", "la"}

	mb = m
	showtype(mb)
	mb.Prints()

	mb = b
	showtype(mb)
	mb.Prints()

	//implicit initialization
	var mb2 Mb = bala{"mu", "gu"}
	showtype(mb2)
	mb2.Prints()

	var mb3 Mb
	//a nil struct bb
	var bb bala
	//a interface with a nil struct
	mb3 = bb
	showtype(mb3)
	mb3.Prints()

	//nil interface
	var mb4 Mb
	showtype(mb4)
	//calling nill interface is a runtime error, because it doesnt know which type of method to call
	//mb4.Prints()

	//empty interface
	var mb5 interface{}
	//holds all types
	showtype(mb5)
	mb5 = "mugu"
	showtype(mb5)
	mb5 = 20
	showtype(mb5)
	mb5 = 10.35
	showtype(mb5)
	mb5 = false
	showtype(mb5)
	mb5 = m
	showtype(mb5)
	mb5 = mb2
	showtype(mb5)
	mb5 = mb3
	showtype(mb5)

	//type assertions
	var mb6 interface{} = "mugu"
	s, ok := mb6.(string)
	if !ok {
		fmt.Println(ok)
	} else {
		fmt.Println(s)
	}

	s1, ok := mb2.(bala)
	if !ok {
		fmt.Println(ok)
	} else {
		fmt.Println(s1)
	}
}

// func showtype(i Mb){
// 	fmt.Printf("the value = %v , the type = %T \n",i,i)
// }

// shows interface value
// a interface holds (value, type) values
// a empty interface may hold value of any type including a interface
func showtype(i interface{}) {
	fmt.Printf("the value = %v , the type = %T \n", i, i)
}
