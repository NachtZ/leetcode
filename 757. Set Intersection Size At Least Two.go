package main

import (
	"fmt"
	"sort"
	"time"
)

type intSlice [][]int

func (s intSlice) Len() int {
	return len(s)
}

func (s intSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s intSlice) Less(i, j int) bool {
	if s[i][0] == s[j][0] {
		return s[j][1] < s[i][1]
	}
	return s[i][0] < s[j][0]
}

func intersectionSizeTwo(intervals [][]int) int {
	sort.Sort(intSlice(intervals))
	todo := make([]int, len(intervals))
	for i := range todo {
		todo[i] = 2
	}
	ans := 0
	for t := len(intervals) - 1; t >= 0; t-- {
		s, m := intervals[t][0], todo[t]
		for p := s; p < s+m; p++ {
			for i := 0; i <= t; i++ {
				if todo[i] > 0 && p <= intervals[i][1] {
					todo[i]--
				}
			}
			ans++
		}

	}
	return ans
}
func main() {
	start := time.Now()
	res := intersectionSizeTwo([][]int{[]int{1, 3}, []int{1, 4}, []int{2, 5}, []int{3, 5}})
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
