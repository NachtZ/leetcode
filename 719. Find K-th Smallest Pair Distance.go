package main

import (
	"fmt"
)

func smallestDistancePair(nums []int, k int) int {
	dis := make([]int, 1000000)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				dis[nums[i]-nums[j]]++
			} else {
				dis[nums[j]-nums[i]]++
			}
		}
	}
	for i := 0; i < len(dis); i++ {
		if k <= dis[i] {
			return i
		}
		k -= dis[i]
	}
	return 0
}
func main() {
	fmt.Println(smallestDistancePair([]int{1, 3, 1}, 2))
}
