package main

import (
	"fmt"
)

/*
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}*/

func minimumDeleteSum(s1 string, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		dp[i] = make([]int, len(s2)+1)
		if i > 0 {
			dp[i][0] = dp[i-1][0] + int(s1[i-1])
		}
	}
	for i := 1; i <= len(s2); i++ {
		dp[0][i] = dp[0][i-1] + int(s2[i-1])
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			tmp := 0
			if s1[i-1] != s2[j-1] {
				tmp = int(s1[i-1]) + int(s2[j-1])
			}
			dp[i][j] = min(dp[i-1][j]+int(s1[i-1]), dp[i][j-1]+int(s2[j-1]))
			dp[i][j] = min(dp[i][j], dp[i-1][j-1]+tmp)
		}
	}
	return dp[len(s1)][len(s2)]

}
func main() {
	fmt.Println(minimumDeleteSum("sea", "eat"))
}
