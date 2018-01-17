package main

import (
	"fmt"
)

/*
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
*/
func maxProfit(prices []int, fee int) int {
	s0, s1 := 0, -65535
	for i := 0; i < len(prices); i++ {
		tmp := s0
		s0 = max(s0, s1+prices[i])
		s1 = max(s1, tmp-prices[i]-fee)
	}
	return s0
}
func main() {
	fmt.Println(maxProfit([]int{2, 1, 4, 4, 2, 3, 2, 5, 1, 2}, 1))
}
