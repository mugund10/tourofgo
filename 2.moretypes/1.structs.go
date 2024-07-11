package main

import "fmt"

type Sex bool

const (
    Male   Sex = true
    Female Sex = false
)

type hooman struct {
    name string
    age  int
    sex  Sex
}

func main() {
    mugu := hooman{"mugu", 24, Male}
    fmt.Println(mugu)
}
