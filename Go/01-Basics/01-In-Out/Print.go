package main

import "fmt"

func main() {
	// One way
	fmt.Println("Golang")

	// Second way
	fmt.Printf("%g\n", 3.2)

	// Third way
	fmt.Print(" ")

	// Some of the most commonly used specifiers are:
	// v – formats the value in a default format.
	// d – formats decimal integers.
	// g – formats the floating-point numbers.
	// b – formats base 2 numbers.
	// o – formats base 8 numbers.
	// t – formats true or false values.
	// s – formats string values.
}
