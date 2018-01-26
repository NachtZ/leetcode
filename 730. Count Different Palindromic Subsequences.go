package main

import (
	"fmt"
)

func countPalindromicSubsequences(S string) int {
	dp := make([][]int, len(S))
	for i := 0; i < len(S); i++ {
		dp[i] = make([]int, len(S))
	}
	mod := 1000000007
	for i := 0; i < len(S); i++ {
		dp[i][i] = 1
	}
	for l := 1; l < len(S); l++ {
		for i := 0; i < len(S)-l; i++ {
			j := i + l
			if S[i] == S[j] {
				left, right := i+1, j-1
				for left <= right && S[left] != S[i] {
					left++
				}
				for left <= right && S[right] != S[i] {
					right--
				}
				if left > right {
					dp[i][j] = (dp[i+1][j-1]*2 + 2) % mod
				} else if left == right {
					dp[i][j] = (dp[i+1][j-1]*2 + 1) % mod
				} else {
					dp[i][j] = (dp[i+1][j-1]*2 - dp[left+1][right-1]) % mod
				}
			} else {
				dp[i][j] = (dp[i][j-1] + dp[i+1][j] - dp[i+1][j-1]) % mod
			}
			if dp[i][j] < 0 {
				dp[i][j] += mod
			}
			dp[i][j] = dp[i][j] % mod
		}
	}
	return dp[0][len(S)-1]

}

func main() {
	fmt.Println(countPalindromicSubsequences("bddaabdbbccdcdcbbdbddccbaaccabbcacbadbdadbccddccdbdbdbdabdbddcccadddaaddbcbcbabdcaccaacabdbdaccbaacc"))
}
