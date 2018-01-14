package main

import (
	"fmt"
)

func hasAlternatingBits(n int) bool {
	before := n % 2
	n /= 2
	for n > 0 {
		if n%2 == before {
			return false
		}
		before = n % 2
		n /= 2
	}
	return true
}

func main() {
	fmt.Println(hasAlternatingBits(9))
}
