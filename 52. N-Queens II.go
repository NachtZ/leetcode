package main

import "fmt"

var ret int

func help(col, left, right []bool, idx, n int) {
	if idx >= n {
		ret++
		return
	} else {
		for i := 0; i < n; i++ {
			if col[i] == true || right[i+idx] == true || left[n-i-1+idx] == true {
				continue
			}
			col[i] = true
			right[i+idx] = true
			left[n-i-1+idx] = true
			help(col, left, right, idx+1, n)
			col[i] = false
			right[i+idx] = false
			left[n-i-1+idx] = false
		}
	}
}

func totalNQueens(n int) int {
	col := make([]bool, n)
	left := make([]bool, n*2-1)
	right := make([]bool, n*2-1)
	ret = 0
	help(col, left, right, 0, n)
	return ret
}
func main() {
	//nums := []int{3, 2, 1}
	fmt.Println(totalNQueens(8))
}
