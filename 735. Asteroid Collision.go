package main

import "fmt"

func asteroidCollision(a []int) []int {
	st := []int{}
	for i := 0; i < len(a); i++ {
		if a[i] > 0 || len(st) == 0 || st[len(st)-1] < 0 {
			st = append(st, a[i])
		} else if st[len(st)-1] <= -a[i] {
			if st[len(st)-1] < -a[i] {
				i--
			}
			st = st[:len(st)-1]
		}
	}
	return st
}
func main() {
	fmt.Println(asteroidCollision([]int{5, 10, -5}))
}
