package main

import (
	"fmt"
)

func isOneBitCharacter(bits []int) bool {
	if len(bits) == 0 {
		return false
	}
	ret := 0
	for i := len(bits) - 2; i >= 0 && bits[i] == 1; i-- {
		ret++
	}
	return ret%2 == 0
}
func main() {
	fmt.Println(isOneBitCharacter([]int{1, 1, 0}))
}
