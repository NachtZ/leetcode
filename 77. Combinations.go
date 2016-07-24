package main

import "fmt"

var ret [][]int

func help(ans []int, n, k int, used []bool) {
	if len(ans) == k {
		tmp := make([]int, k)
		copy(tmp, ans)
		ret = append(ret, tmp)
		return
	}
	i := 1
	if len(ans) > 0 {
		i = ans[len(ans)-1] + 1
	}
	for ; i <= n; i++ {
		if used[i] == true {
			continue
		}
		used[i] = true
		ans = append(ans, i)
		help(ans, n, k, used)
		ans = ans[:len(ans)-1]
		used[i] = false
	}
}
func combine(n int, k int) [][]int {
	ret = [][]int{}
	used := make([]bool, n+1)
	help([]int{}, n, k, used)
	return ret
}
func main() {
	//fmt.Println(isNumber("123456"))

	fmt.Println(combine(4, 2))
	fmt.Println("Hello")
}
