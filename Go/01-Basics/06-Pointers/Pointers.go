package main

import "fmt"

func main() {
	i, j := 1, 2
	fmt.Println(i, j)
	fmt.Println(&i, &j)
	k, l := &i, &j
	fmt.Println(k, l)
	fmt.Println(*k, *l)
	swap(&i, &j)
	fmt.Println(i, j)
}
func swap(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp

}
