package main

import (
	"fmt"
	"sort"
)

func helper(nums []int, v []int, l, sum, k, target int) bool {
	if k == 0 {
		return true
	}
	if sum == target {
		return helper(nums, v, 0, 0, k-1, target)
	}
	if sum > target {
		return false
	}
	for i := l; i < len(nums); i++ {
		if v[i] == 1 {
			continue
		}
		v[i] = 1
		if helper(nums, v, i+1, sum+nums[i], k, target) {
			return true
		}
		v[i] = 0
	}
	return false
}

func canPartitionKSubsets(nums []int, k int) bool {
	if k == 1 {
		return true
	}
	sort.Ints(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k > 0 {
		return false
	}
	v := make([]int, len(nums))
	return helper(nums, v, 0, 0, k, sum/k)
}
func main() {
	fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
}
