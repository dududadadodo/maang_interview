package main

import "fmt"

func maxArea(heights []int) int {
	max_area := 0
	for i := 0; i < len(heights); i++ {
		for k := i + 1; k < len(heights); k++ {
			length := min(heights[i], heights[k])
			width := k - i
			area := length * width
			max_area = max(max_area, area)
		}
	}
	return max_area
}
func main() {
	fmt.Print(maxArea([]int{5, 9, 2, 4}))
}
