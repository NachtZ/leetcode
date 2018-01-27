package main

import (
	"fmt"
	"time"
)

func dominantIndex(nums []int) int {
	m1, m2 := 0, 1
	if len(nums) <= 1 {
		return 0
	}
	if nums[0] < nums[1] {
		m1, m2 = 1, 0
	}
	for i := 2; i < len(nums); i++ {
		if nums[i] > nums[m1] {
			m1, m2 = i, m1
		} else if nums[i] > nums[m2] {
			m2 = i
		}
	}
	if nums[m2]*2 <= nums[m1] {
		return m1
	}
	return -1
}
func main() {
	start := time.Now()
	res := dominantIndex([]int{3, 6, 1, 0})
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
