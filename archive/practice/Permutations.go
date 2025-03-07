package main

import "fmt"

func main() {
	var up string
	fmt.Print("String: ")
	fmt.Scan(&up)
	Permutations("", up)
}
func Permutations(p, up string) {
	if up == "" {
		fmt.Println(p)
		return
	}
	var ch = up[0:1]
	Answer(p, up, ch, 0)
}
func Answer(p, up, ch string, r int) {
	if r > len(p) {
		return
	}
	var first = p[0:r]
	var second = p[r:len(p)]
	Permutations(first+ch+second, up[1:])
	Answer(p, up, ch, r+1)
}
