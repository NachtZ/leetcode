package main

import (
	"fmt"
)

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	sums := make([]uint64, len(nums)+1)
	sums[0] = 0
	left, right := make([]int, len(nums)), make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sums[i+1] = uint64(nums[i]) + sums[i]
		right[i] = len(nums) - k
	}

	for i, s := k, sums[k]-sums[0]; i < len(nums); i++ {
		if sums[i+1]-sums[i+1-k] > s {
			left[i] = i + 1 - k
			s = sums[i+1] - sums[i+1-k]
		} else {
			left[i] = left[i-1]
		}
	}
	for i, s := len(nums)-1-k, sums[len(nums)]-sums[len(nums)-k]; i >= 0; i-- {
		if sums[i+k]-sums[i] >= s {
			right[i] = i
			s = sums[i+k] - sums[i]
		} else {
			right[i] = right[i+1]
		}
	}
	res, ret := uint64(0), []int{}
	for i := k; i <= len(nums)-2*k; i++ {
		l, r := left[i-1], right[i+k]
		s := sums[i+k] - sums[i] + sums[l+k] - sums[l] + sums[r+k] - sums[r]
		if s > res {
			res = s
			ret = []int{l, i, r}
		}
	}
	return ret
}

func main() {
	fmt.Println(maxSumOfThreeSubarrays([]int{1, 2, 1, 2, 6, 7, 5, 1}, 2))
}
