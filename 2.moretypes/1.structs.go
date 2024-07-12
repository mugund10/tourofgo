package main

import (
    "fmt"
    "github.com/mugund10/tourofgo/2.moretypes/packs"
)

var mugu bool = true

func main() {
    hooman := packs.Hooman{}
    mugu = hooman.NewHooman("mugu", 24, "Male")
    fmt.Println(mugu)
}
