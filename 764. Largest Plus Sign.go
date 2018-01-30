package main

import (
	"fmt"
	"time"
)

func orderOfLargestPlusSign(N int, mines [][]int) int {
	m, dp := make(map[int]bool), make([][]int, N)
	for _, mine := range mines {
		m[mine[0]*N+mine[1]] = true
	}
	for i := range dp {
		dp[i] = make([]int, N)
	}
	ans, cnt := 0, 0
	for r := 0; r < N; r++ {
		cnt = 0
		for c := 0; c < N; c++ {
			if m[r*N+c] {
				cnt = 0
			} else {
				cnt++
			}
			dp[r][c] = cnt
		}
		cnt = 0
		for c := N - 1; c >= 0; c-- {
			if m[r*N+c] {
				cnt = 0
			} else {
				cnt++
			}
			if cnt < dp[r][c] {
				dp[r][c] = cnt
			}
		}
	}
	for c := 0; c < N; c++ {
		cnt = 0
		for r := 0; r < N; r++ {
			if m[r*N+c] {
				cnt = 0
			} else {
				cnt++
			}
			if cnt < dp[r][c] {
				dp[r][c] = cnt
			}
		}
		cnt = 0
		for r := N - 1; r >= 0; r-- {
			if m[r*N+c] {
				cnt = 0
			} else {
				cnt++
			}
			if dp[r][c] > cnt {
				dp[r][c] = cnt
			}
			if ans < dp[r][c] {
				ans = dp[r][c]
			}
		}
	}
	return ans
}
func main() {
	start := time.Now()
	res := orderOfLargestPlusSign(5, [][]int{[]int{4, 2}})
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
