package main

import "fmt"

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		if dp[i] == false {
			continue
		}
		for j := 1; j <= nums[i] && j+i < len(nums); j++ {
			dp[i+j] = true
		}
	}
	return dp[len(nums)-1]
}
func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
