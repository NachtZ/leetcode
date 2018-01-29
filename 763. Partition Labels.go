package main

import (
	"fmt"
	"sort"
	"time"
)

type MyInterval struct {
	Start int
	End   int
}

type IntervalSlice []MyInterval

func (s IntervalSlice) Less(i, j int) bool {
	if s[i].Start != s[j].Start {
		return s[i].Start < s[j].Start
	}
	return s[i].End < s[j].End
}
func (s IntervalSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s IntervalSlice) Len() int {
	return len(s)
}

func partitionLabels(S string) []int {
	m := make(map[rune]MyInterval)
	for i, b := range S {
		if _, ok := m[b]; ok {
			m[b] = MyInterval{m[b].Start, i}
		} else {
			m[b] = MyInterval{i, i}
		}
	}
	s := IntervalSlice{}
	for _, v := range m {
		s = append(s, v)
	}
	sort.Sort(s)
	//fmt.Println(s)
	ret, start, end := []int{}, s[0].Start, s[0].End
	for i := 1; i < len(s); i++ {
		if s[i].Start > end {
			ret = append(ret, end-start+1)
			start, end = s[i].Start, s[i].End
		} else if s[i].End > end {
			end = s[i].End
		}
	}
	ret = append(ret, end-start+1)
	return ret
}
func main() {
	start := time.Now()
	res := partitionLabels("ababcbacadefegdehijhklij")
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
