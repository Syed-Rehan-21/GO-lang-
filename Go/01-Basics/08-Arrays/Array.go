package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var b [3]int
	for i := 0; i < 3; i++ {
		fmt.Scan(&b[i])
	}
	for i := 0; i < 3; i++ {
		fmt.Print(b[i], " ")
	}
}
