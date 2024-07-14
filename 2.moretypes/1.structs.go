package main

import "fmt"

type cname string
type rupees float64
type released int

type car struct {
	name  cname
	price rupees
	year  released
}

//to make this work we need to return values to the main function
func (c car) update(a ...any) {

	for _, v := range a {
		//the v from the for and swith is different but being same type
		//refer
		switch v := v.(type) {
		case cname:
			c.name = v
		case rupees:
			c.price = v
		case released:
			c.year = v
		}

	}

}

func (c *car) updatePOINTER(a ...any) {

	for _, v := range a {
		//the v from the for and swith is different but being same type
		//refer
		switch v := v.(type) {
		case cname:
			c.name = v
		case rupees:
			c.price = v
		case released:
			c.year = v
		}

	}

}

func (c car) updateRETURN(a ...any) any{

	for _, v := range a {
		//the v from the for and swith is different but being same type
		//its a type switch
		switch v := v.(type) {
		case cname:
			c.name = v
		case rupees:
			c.price = v
		case released:
			c.year = v
		}

	}
    return c
}

func main() {

	oldcar := car{cname("mustaaang"), rupees(740000), released(2023)}
	fmt.Println(oldcar)

	oldcar.update(cname("mustaang"), released(2024))
	fmt.Println("func update : ",oldcar)  

    oldcar.updatePOINTER(cname("mustaang"), released(2024))
	fmt.Println("func updatePOINTER : ",oldcar)

    oldcarupdate := oldcar.updateRETURN(cname("mustang"), released(2024))
	fmt.Println("func updateRETURN : ",oldcarupdate)

	newcar := car{cname("gowagon"), rupees(4200000), released(2019)}
    fmt.Println(newcar)

    newcar.name = "G wagan"
    fmt.Println(newcar)

}
