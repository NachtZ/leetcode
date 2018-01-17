package main

import (
	"fmt"
)

/*
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
*/
func findLength(A []int, B []int) int {
	dp := make([][]int, len(A)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(B)+1)
	}
	ret := 0
	for i := len(A) - 1; i >= 0; i-- {
		for j := len(B) - 1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}
			ret = max(dp[i][j], ret)
		}
	}
	return ret
}
func main() {
	fmt.Println(findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
}
