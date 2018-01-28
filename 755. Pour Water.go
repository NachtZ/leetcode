package main

import (
	"fmt"
	"time"
)

func pourWater(heights []int, V int, K int) []int {
	for ; V > 0; V-- {
		drop := true
		for d := -1; d <= 1; d += 2 {
			i, best := K, K
			for i+d >= 0 && i+d < len(heights) && heights[i+d] <= heights[i] {
				if heights[i+d] < heights[i] {
					best = i + d
				}
				i += d
			}
			if heights[best] < heights[K] {
				heights[best]++
				drop = false
				break
			}
		}
		if drop {
			heights[K]++
		}
	}
	return heights
}
func main() {
	start := time.Now()
	res := pourWater([]int{2, 1, 1, 2, 1, 2, 2}, 4, 3)
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
