package main

import (
	"fmt"
	"time"
)

func isIdealPermutation(A []int) bool {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	for i, v := range A {
		if abs(v-i) > 1 {
			return false
		}
	}
	return true
}
func main() {
	start := time.Now()
	res := isIdealPermutation([]int{1, 0, 2})
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
