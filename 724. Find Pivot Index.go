package main

import (
	"fmt"
)

func pivotIndex(nums []int) int {
	if len(nums) <= 1 {
		return len(nums) - 1
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum == nums[0] {
		return 0
	}
	before := 0
	for i := 1; i < len(nums); i++ {
		before += nums[i-1]
		if sum-nums[i]-before == before {
			return i
		}
	}
	return -1
}
func main() {
	fmt.Println(pivotIndex([]int{-1, -1, -1, -1, 1, 1}))
}
