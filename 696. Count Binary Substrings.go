package main

import (
	"fmt"
)

func countBinarySubstrings(s string) int {
	p, c := 0, 1
	cnt := 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			c++
		} else {
			p = c
			c = 1
		}
		if p >= c {
			cnt++
		}

	}
	return cnt
}
func main() {
	fmt.Println(countBinarySubstrings("10101"))
}
