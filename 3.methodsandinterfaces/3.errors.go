package main

import (
	"fmt"
)

type str string

func main() {
	a := str("a")
	b := str("")
	now, err := kumkum(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(now)
	}
}

func kumkum(a, b str) (string, error) {
	if string(a) == "" {
		return "", &a
	} else if string(b) == "" {
		return "", &b
	} else {
		return fmt.Sprintf("%s%s", a, b), nil
	}
}

func (s *str) Error() string {
	return fmt.Sprintf("input string is empty")
}
