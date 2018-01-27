package main

import (
	"fmt"
	"time"
)

func cherryPickup(grid [][]int) int {
	mx := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n, maxp := len(grid), 2*len(grid)-1
	dp := make([][]int, len(grid))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(grid))
		for j := 0; j < len(dp); j++ {
			dp[i][j] = -1
		}
	}
	dp[0][0] = grid[0][0]
	for p := 1; p < maxp; p++ {
		for i := n - 1; i >= 0; i-- {
			for x := n - 1; x >= 0; x-- {
				j, y := p-i, p-x
				if j < 0 || j >= n || y < 0 || y >= n || grid[x][y] == -1 || grid[i][j] == -1 {
					dp[i][x] = -1
					continue
				}
				if i > 0 {
					dp[i][x] = mx(dp[i][x], dp[i-1][x])
				}
				if x > 0 {
					dp[i][x] = mx(dp[i][x], dp[i][x-1])
				}
				if i > 0 && x > 0 {
					dp[i][x] = mx(dp[i][x], dp[i-1][x-1])
				}
				if dp[i][x] >= 0 {
					dp[i][x] += grid[i][j]
					if i != x {
						dp[i][x] += grid[x][y]
					}
				}
			}
		}
	}
	return mx(dp[n-1][n-1], 0)
}
func main() {
	start := time.Now()
	res := cherryPickup([][]int{[]int{0, 1, -1}, []int{1, 0, -1}, []int{1, 1, 1}})
	end := time.Now()
	fmt.Println(start, end)
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
