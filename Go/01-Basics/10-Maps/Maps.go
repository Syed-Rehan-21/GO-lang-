package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answerr"] = 48
	fmt.Println("The value:", m["Answerr"])

	fmt.Println("map:", m)

	v1 := m["Answerr"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "Answerr")
	fmt.Println("map:", m)
	v, ok := m["Answerr"]
	fmt.Println("The value:", v, "Present?", ok)
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
