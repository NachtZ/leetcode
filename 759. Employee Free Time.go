package main

import (
	"fmt"
	"sort"
	"time"
)

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

type IntervalSlice []Interval

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

func employeeFreeTime(schedule [][]Interval) []Interval {
	all := []Interval{}
	for _, s := range schedule {
		for _, sc := range s {
			all = append(all, sc)
		}
	}
	if len(all) == 0 {
		return []Interval{}
	}
	sort.Sort(IntervalSlice(all))
	ret, start := []Interval{}, all[0].End
	for i := 1; i < len(all); i++ {
		if all[i].Start > start {
			ret, start = append(ret, Interval{start, all[i].Start}), all[i].End
		} else if all[i].End >= start {
			start = all[i].End
		}
	}
	//ret = append(ret, Interval{start, end})
	return ret
}
func main() {
	start := time.Now()
	res := employeeFreeTime([][]Interval{[]Interval{Interval{1, 2}, Interval{5, 6}}, []Interval{Interval{1, 3}, Interval{4, 10}}})
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
