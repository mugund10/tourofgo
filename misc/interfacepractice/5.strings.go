package main

import (
	"fmt"

	"github.com/mugund10/tourofgo/misc/interfacepractice"
)

// Define the mugus type
type mugus string

// Implement the Kuruma method for the mugus type
func (m mugus) Kuruma() string {
	return fmt.Sprintf("ans = %v", m)
}

func main() {
	fmt.Println("a")
	interfacepractice.Podu("A")

	// Create an instance of mugus
	abd := mugus("kuruma")

	// Call the Kuruma method on the mugus instance and pass it to Podu
	interfacepractice.Podu(abd)
}
