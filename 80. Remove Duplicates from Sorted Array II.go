package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	ptr := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] == nums[ptr-2] {
			continue
		}
		nums[ptr] = nums[i]
		ptr++
	}
	return ptr
}
func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1}))
	fmt.Println("Hello")
}
