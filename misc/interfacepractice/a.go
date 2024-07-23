package interfacepractice

import "fmt"

type Kuruma interface {
	Kuruma() string
}

func Podu(a any) {
	if k, ok := a.(Kuruma) ; ok {
		fmt.Println(k.Kuruma())
	}else{
	fmt.Println(a)
	}
}
