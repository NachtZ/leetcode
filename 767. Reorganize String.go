package main

import (
	"fmt"
	"sort"
	"time"
)

func reorganizeString(S string) string {
	m := make([]int, 26)
	if len(S) <= 1 {
		return S
	}
	for _, s := range S {
		m[s-'a'] += 100
	}
	for i, v := range m {
		m[i] = v + i
	}
	sort.Ints(m)
	ret := make([]rune, len(S))
	t := 1
	for _, c := range m {
		cnt, ch := c/100, rune(c%100+'a')
		if cnt > (len(S)+1)/2 {
			return ""
		}
		for i := 0; i < cnt; i++ {
			if t >= len(S) {
				t = 0
			}
			ret[t] = ch
			t += 2

		}
	}
	return string(ret)

}

/*
[[11,74,0,93],
 [40,11,74,7]]
*/
func main() {
	start := time.Now()
	res := reorganizeString("aab")
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
