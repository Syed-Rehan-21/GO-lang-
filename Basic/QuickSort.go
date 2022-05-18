package main

import "fmt"

func main() {
	var n int
	fmt.Print("Size of the array: ")
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Print("Array elements: ")
	Input(0, n, arr)
	Sort(0, len(arr)-1, arr)
	Output(arr)
}
func Input(r, size int, arr []int) {
	if r > size-1 {
		return
	}
	fmt.Scan(&arr[r])
	Input(r+1, size, arr)
}
func Sort(low, high int, arr []int) {
	if low < high {
		pivotIndex := Partition(low, high, arr)
		Sort(low, pivotIndex-1, arr)
		Sort(pivotIndex+1, high, arr)
	}
}
func Partition(low, high int, arr []int) int {
	pivot := arr[high] // Let, pivot be the last element
	i := low - 1
	index := Arranging(pivot, i, low, high, arr)
	Swap(index+1, high, arr)
	return index + 1
}
func Arranging(pivot, i, r, high int, arr []int) int {
	if r > high-1 {
		return i
	}
	if arr[r] < pivot {
		i++
		Swap(i, r, arr)
		return Arranging(pivot, i, r+1, high, arr)
	}
	return Arranging(pivot, i, r+1, high, arr)
}
func Swap(first, second int, arr []int) {
	temp := arr[first]
	arr[first] = arr[second]
	arr[second] = temp
}
func Output(arr []int) {
	fmt.Println("After Sorting :", arr)
}
