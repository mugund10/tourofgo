package main

import "fmt"

type in int

type box struct {
	x, y int
}

type somestep interface {
	add(a... int) int
	mult(a... int) int
}

func main() {
	a := in(85)
	b := box{2, 3}
	var c somestep
	c = &a
	c.add(2,3,4,5)
	c.mult(1,2,3)
	showtypes(c)
	c = &b
	c.add(1,2,3,5)
	c.mult(2,5)
	showtypes(c)
}

func (i *in) add(a... int) int {
	for it := range a {
		*i += in(a[it])
	}
	
	fmt.Println("add in =", *i)
	return int(*i)
}

func (i *in) mult(a... int) int {
	for it := range a{
		*i *= in(a[it])
	}
	fmt.Println("mult in =", *i)
	return int(*i)
}

func (b *box) add(a... int) int {
	for it := range a{
		if it % 2 == 1{
			b.y += a[it]
		}else {
			b.x += a[it]
		}
	}
	// b.x += 1
	// b.y += 2
	fmt.Println("add box =", *b)
	return int(b.x + b.y)
}

func (b *box) mult(a... int) int {
	for it := range a {
		if it % 2 == 1 {
			b.y *= a[it]
		}else {
			b.x *= a[it]
		}
	}
	fmt.Println("mult box =", *b)
	return int(b.x * b.y)
}

func showtypes(i interface{}) {
	switch v := i.(type) {
	case *in:
		fmt.Printf("the value = &%v , the type = %T\n", *v, v)
	case *box:
		fmt.Printf("the value = &%v , the type = %T\n", *v, v)
	default:
		fmt.Printf("the value = %v , the type = %T\n", v, v)
	}
}
