package main

import "fmt"

type ER string

func (s ER) Error() string {
	return fmt.Sprintf("the error is %v", string(s))
}

func schedules(s string)error{
	return fmt.Errorf(s)
}

func schedule(s string) (error) {
	if s == "err" {
		return  ER(s)
	} else {
		return  nil
	}

}

func main() {
	fmt.Println(schedules("mugu"))
	fmt.Println(schedule("err"))

}
