package main

import (
	"fmt"
	"time"
)

func countPrimeSetBits(L int, R int) int {
	cntByte := func(x int) int {
		cnt := 0
		for x > 0 {
			x &= (x - 1)
			cnt++
		}
		return cnt
	}
	cnt := 0
	for i := L; i <= R; i++ {
		x := cntByte(i)
		if x == 2 || x == 3 || x == 5 || x == 7 ||
			x == 11 || x == 13 || x == 17 || x == 19 {
			cnt++
		}
	}
	return cnt
}
func main() {
	start := time.Now()
	res := countPrimeSetBits(1, 100)
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
