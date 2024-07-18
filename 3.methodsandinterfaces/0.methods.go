package main

import "fmt"

type gov struct {
	x, y float64
}

// method
func (g gov) salary() float64 {
	return g.x + g.y
}

// method as function
func salary(g gov) float64 {
	return g.x + g.y
}

// pointer receivers method
func (g *gov) salarypointer(f float64) {
	g.x += f
	g.y += f
}

// pointer receivers as func
func salarypointer(g *gov, f float64) {
	g.x += f
	g.y += f
}

type money float64

// we can declare a method on non struct type too
func (m money) printMoney() {
	fmt.Println("the value is : ", m)
}

func main() {
	jaya := gov{60000, 40000}
	// method
	fmt.Println(jaya.salary())
	fmt.Println(jaya)
	// method as func
	fmt.Println(salary(jaya))
	fmt.Println(jaya)

	// pointer receivers method
	(&jaya).salarypointer(100)
	// pointer indirection
	jaya.salarypointer(100)
	fmt.Println(jaya.salary())

	// pointer receiver as func
	salarypointer(&jaya, 100)
	fmt.Println(jaya.salary())

	inr := money(100.0)

	inr.printMoney()

}
