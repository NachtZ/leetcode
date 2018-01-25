package main

import (
	"fmt"
)

func selfDividingNumbers(left int, right int) []int {
	ret := []int{}
	for i := left; i <= right; i++ {
		flag := true
		for tmp := i; flag && tmp > 0; {
			if tmp%10 == 0 || i%(tmp%10) > 0 {
				flag = false
			}
			tmp /= 10
		}
		if flag {
			ret = append(ret, i)
		}
	}
	return ret
}
func main() {
	fmt.Println(selfDividingNumbers(1, 10000))
}
