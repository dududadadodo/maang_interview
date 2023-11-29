package main

import "fmt"

func maxArea(heights []int) int {
	max_area := 0
	l := 0
	r := len(heights) - 1
	for l < r {
		length := min(heights[l], heights[r])
		width := r - l
		area := length * width
		max_area = max(max_area, area)

		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}
	return max_area
}
func main() {
	fmt.Print(maxArea([]int{5, 9, 2, 4}))
}
