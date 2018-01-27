package main

import "fmt"

//got TLE, may be leetcode's bug.
func deleteAndEarn(nums []int) int {
	num := make([]int, 10001)
	selectbefore, skipbefore := 0, 0
	maxv, minv := 0, 10001
	for _, v := range nums {
		num[v] += v
		if v > maxv {
			maxv = v
		}
		if v < minv {
			minv = v
		}
	}
	for i := minv; i <= maxv; i++ {
		selectbefore, skipbefore = skipbefore+num[i], func() int {
			if selectbefore > skipbefore {
				return selectbefore
			}
			return skipbefore
		}()
	}
	if selectbefore > skipbefore {
		return selectbefore
	}
	return skipbefore
}
func main() {
	res := deleteAndEarn([]int{2, 3, 4})
	fmt.Println(res)
}
