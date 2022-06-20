package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Two Numbers Please: ")
	fmt.Scan(&a, &b)
	fmt.Print(a + b)
}
