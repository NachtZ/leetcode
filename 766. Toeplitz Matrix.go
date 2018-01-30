package main

import (
	"fmt"
	"time"
)

func isToeplitzMatrix(matrix [][]int) bool {
	for i := 0; i < len(matrix); i++ {
		x, y := i, 0
		for j := 1; true; j++ {
			x, y = x+1, y+1
			if x < 0 || y < 0 || x >= len(matrix) || y >= len(matrix[0]) {
				break
			}

			if matrix[x][y] != matrix[x-1][y-1] {
				return false
			}
		}
	}
	for i := 1; i < len(matrix[0]); i++ {
		x, y := 0, i
		for j := 1; true; j++ {
			x, y = x+1, y+1
			if x < 0 || y < 0 || x >= len(matrix) || y >= len(matrix[0]) {
				break
			}
			if matrix[x][y] != matrix[x-1][y-1] {
				return false
			}
		}
	}
	return true
}

/*
[[11,74,0,93],
 [40,11,74,7]]
*/
func main() {
	start := time.Now()
	res := isToeplitzMatrix([][]int{[]int{11, 74, 0, 93}, []int{40, 11, 74, 7}})
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
