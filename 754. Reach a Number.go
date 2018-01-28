package main

import (
	"fmt"
	"time"
)

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	k := 0
	for target > 0 {
		k++
		target -= k
	}
	if target%2 == 0 {
		return k
	}
	return k + 1 + k%2
}
func main() {
	start := time.Now()
	res := reachNumber(3)
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
