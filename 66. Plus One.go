package main

import "fmt"

func plusOne(digits []int) []int {
	plus := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += plus
		if digits[i] >= 10 {
			plus = 1
			digits[i] -= 10
		} else {
			return digits
		}
	}
	if plus == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	//fmt.Println(isNumber("123456"))
	fmt.Println(plusOne([]int{9}))
	fmt.Println("Hello")
}
