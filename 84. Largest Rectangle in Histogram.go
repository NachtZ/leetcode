package main

import "fmt"

func largestRectangleArea(heights []int) int {
	stack := []int{}
	heights = append(heights, 0)
	max := 0
	for i := 0; i < len(heights); i++ {
		if len(stack) == 0 || heights[i] > heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				tmp = heights[tmp] * i
			} else {
				tmp = heights[tmp] * (i - stack[len(stack)-1] - 1)
			}
			if tmp > max {
				max = tmp
			}
			i--
		}
	}
	return max
}
func main() {
	fmt.Println(largestRectangleArea([]int{1}))
	fmt.Println("Hello")
}
