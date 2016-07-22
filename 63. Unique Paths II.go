package main

func uniquePathsWithObstacles(dp [][]int) int {

	m := len(dp)
	n := len(dp[0])
	if dp[0][0] == 1 {
		return 0
	}
	dp[0][0] = 1
	for i := 1; i < m; i++ {
		if dp[i][0] != 1 {
			dp[i][0] = dp[i-1][0]
		} else {
			dp[i][0] = 0
		}
	}
	for i := 1; i < n; i++ {
		if dp[0][i] != 1 {
			dp[0][i] = dp[0][i-1]
		} else {
			dp[0][i] = 0
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else {
				dp[i][j] = 0
			}
		}
	}
	return dp[m-1][n-1]
}
