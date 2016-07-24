package main

import "fmt"

func mySqrt(x int) int {
	if x == 0 {
		return x
	}
	if x < 4 {
		return 1
	}
	left := 2
	right := int(x / 2)
	for left < right {
		if left == right-1 {
			return left
		}
		mid := int((left + right) / 2)
		res := mid * mid
		if res == x {
			return mid
		}
		if res < x {
			left = mid
		} else {
			right = mid
		}
	}
	return left

}
func main() {
	//fmt.Println(isNumber("123456"))

	fmt.Println(mySqrt(99))
	fmt.Println("Hello")
}
