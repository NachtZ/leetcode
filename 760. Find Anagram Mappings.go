package main

import (
	"fmt"
	"time"
)

func anagramMappings(A []int, B []int) []int {
	m := make(map[int]int)
	for i, v := range B {
		m[v] = i
	}
	ret := []int{}
	for _, v := range A {
		ret = append(ret, m[v])
	}
	return ret
}
func main() {
	start := time.Now()
	res := anagramMappings([]int{12, 28, 46, 32, 50}, []int{50, 12, 32, 46, 28})
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
