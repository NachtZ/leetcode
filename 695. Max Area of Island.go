package main

import (
	"fmt"
)

func travel(grid [][]int, v [][]int, x, y int) int {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || v[x][y] == 1 || grid[x][y] == 0 {
		return 0
	}
	moves := [][]int{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}}
	v[x][y] = 1
	ret := 1
	for i := 0; i < 4; i++ {
		ret += travel(grid, v, x+moves[i][0], y+moves[i][1])
	}
	return ret
}

func maxAreaOfIsland(grid [][]int) int {
	v := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		v[i] = make([]int, len(grid[0]))
	}
	ret := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if v[i][j] == 1 || grid[i][j] == 0 {
				continue
			}
			ans := travel(grid, v, i, j)
			if ans > ret {
				ret = ans
			}
		}
	}
	return ret
}
func main() {
	fmt.Println(maxAreaOfIsland([][]int{[]int{0, 0, 0, 0, 0, 0, 0, 0}}))
}
