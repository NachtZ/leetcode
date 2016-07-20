package main

import "fmt"

var ret [][]string

func help(b []int, col, left, right []bool, idx, n int) {
	if idx >= n {
		tmp := []string{}
		for i := 0; i < n; i++ {
			str := ""
			for j := 0; j < b[i]; j++ {
				str += "."
			}
			str += "Q"
			for j := b[i] + 1; j < n; j++ {
				str += "."
			}
			tmp = append(tmp, str)
		}
		ret = append(ret, tmp)
	} else {
		for i := 0; i < n; i++ {
			if col[i] == true || right[i+idx] == true || left[n-i-1+idx] == true {
				continue
			}
			col[i] = true
			right[i+idx] = true
			left[n-i-1+idx] = true
			b[idx] = i
			help(b, col, left, right, idx+1, n)
			col[i] = false
			right[i+idx] = false
			left[n-i-1+idx] = false
		}
	}
}

func solveNQueens(n int) [][]string {
	col := make([]bool, n)
	left := make([]bool, n*2-1)
	right := make([]bool, n*2-1)
	board := make([]int, n)
	ret = [][]string{}
	help(board, col, left, right, 0, n)
	return ret
}
func main() {
	//nums := []int{3, 2, 1}
	fmt.Println(solveNQueens(4))
}
