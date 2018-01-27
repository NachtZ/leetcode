package main

import (
	"fmt"
	"time"
)

func minCostClimbingStairs(cost []int) int {
	if len(cost) < 2 {
		return 0
	}
	dp := make([]int, len(cost)+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i <= len(cost); i++ {
		dp[i] = cost[i-1] + dp[i-1]
		if dp[i] > cost[i-2]+dp[i-2] {
			dp[i] = cost[i-2] + dp[i-2]
		}
	}
	return dp[len(cost)]

}
func main() {
	start := time.Now()
	res := minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
