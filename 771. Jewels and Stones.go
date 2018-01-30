package main

import (
	"fmt"
	"time"
)

func numJewelsInStones(J string, S string) int {
	m := make(map[rune]bool)
	for _, v := range J {
		m[v] = true
	}
	cnt := 0
	for _, v := range S {
		if m[v] == true {
			cnt++
		}
	}
	return cnt
}

func main() {
	start := time.Now()
	res := numJewelsInStones("Aa", "AaaaaaaBBB")
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
