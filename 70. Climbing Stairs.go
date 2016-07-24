package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}
func main() {
	//fmt.Println(isNumber("123456"))

	fmt.Println(climbStairs(1))
	fmt.Println("Hello")
}
