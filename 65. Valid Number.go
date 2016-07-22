package main

import "fmt"

func isNumber(s string) bool {
	table := [][]int{{-1, 1, -1, 0, 3, 4}, {-1, 1, 2, 5, 10, -1}, {-1, 6, -1, -1, -1, 9},
		{-1, 7, -1, -1, -1, -1}, {-1, 8, -1, -1, 3, -1}, {-1, -1, -1, 5, -1, -1}, {-1, 6, -1, 5, -1, -1},
		{-1, 7, 2, 5, -1, -1}, {-1, 8, 2, 5, 10, -1}, {-1, 6, -1, -1, -1, -1}, {-1, 7, 2, 5, -1, -1}}
	state := 0
	exit := []bool{false, true, false, false, false, true, true, true, true, false, true}
	var charmap [256]int
	for i := int('0'); i <= int('9'); i++ {
		charmap[i] = 1
	}
	charmap[int('e')] = 2
	charmap[int(' ')] = 3
	charmap[int('.')] = 4
	charmap[int('+')] = 5
	charmap[int('-')] = 5
	for _, char := range s {
		state = table[state][charmap[int(char)]]
		if state == -1 {
			return false
		}
	}
	return exit[state]
}

func main() {
	fmt.Println(isNumber("123456"))
}
