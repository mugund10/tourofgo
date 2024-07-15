package main

import "fmt"

func main() {
	sli := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sli)
	fmt.Printf("len : %d , cap : %d \n", len(sli), cap(sli))

	sli = sli[2:5]
	fmt.Println("sli[2:5]")
	fmt.Println(sli)
	fmt.Printf("len : %d , cap : %d \n", len(sli), cap(sli))

	sli = sli[3:5]
	fmt.Println("sli[3:5]")
	fmt.Println(sli)
	fmt.Printf("len : %d , cap : %d \n", len(sli), cap(sli))

	sli = sli[:5]
	fmt.Println("sli[:5]")
	fmt.Println(sli)
	fmt.Printf("len : %d , cap : %d \n", len(sli), cap(sli))

	// creating slice using make()
	sli2 := make([]int, 0, 5)
	fmt.Println("make([]int,2,5)")
	fmt.Println(sli2)
	fmt.Printf("len : %d , cap : %d \n", len(sli2), cap(sli2))

	//adding values to the slice using append()

	sli2 = append(sli2, 10, 12, 14, 16, 18, 20)
	fmt.Println("append(sli2, 10,12,14,16)")
	fmt.Println(sli2)
	fmt.Printf("len : %d , cap : %d \n", len(sli2), cap(sli2))

}
