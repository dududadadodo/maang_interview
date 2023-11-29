package main

import "fmt"

func Calc(array []int) bool {
	i := 1
	for i < len(array) && array[i] > array[i-1] {
		i++
	}

	if i == 1 || i == len(array) {
		return false
	}
	for i < len(array) && array[i] < array[i-1] {
		i++
	}

	return i == len(array)
}
func main() {
	fmt.Print(Calc([]int{0, 2, 3, 4, 5, 2, 1}))
}
