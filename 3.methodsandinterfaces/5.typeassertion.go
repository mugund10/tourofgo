package main

import "fmt"

type ak47 struct {
}

type m416 struct {
}

func (ak47) shoot() {
	fmt.Println("tak tak tak")
}
func (m416) shoot() {
	fmt.Println("thop thop thop")
}

func (m416) addscope() {
	fmt.Println("cling")
}

type guns interface {
	shoot()
}

func main() {

	rg := []guns{ak47{}, m416{}}

	for _, val := range rg {
		val.shoot()
		if re, ok := val.(m416); ok {
				re.addscope()
		}
	}

}
