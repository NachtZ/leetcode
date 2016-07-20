package main

import "fmt"

func spiralOrder(mx [][]int) []int {
	n := len(mx)
	if n == 0 {
		return []int{}
	}
	if n <= 1 {
		return mx[0]
	}
	m := len(mx[0])
	ret := make([]int, m*n)
	idx, max := 0, m*n
	for i := 0; idx < max; i++ {
		for j := i; j < m-i && idx < max; j++ {
			ret[idx] = mx[i][j]
			idx++
		}
		for j := i + 1; j < n-i && idx < max; j++ {
			ret[idx] = mx[j][m-i-1]
			idx++
		}
		for j := m - i - 2; j >= i && idx < max; j-- {
			ret[idx] = mx[n-i-1][j]
			idx++
		}
		for j := n - i - 2; j > i && idx < max; j-- {
			ret[idx] = mx[j][i]
			idx++
		}
	}
	return ret
}

func main() {
	m := [][]int{}
	m = append(m, []int{1, 2, 3})

	m = append(m, []int{4, 5, 6})
	m = append(m, []int{7, 8, 9})
	m = append(m, []int{10, 11, 12})
	//m = append(m, []int{7, 8, 9})
	fmt.Println(spiralOrder(m))
}
