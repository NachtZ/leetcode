package main

import "fmt"

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	ml, mr, min := 0, len(s)-1, -1*len(s)
	l, r := 0, mr
	var cmap [260]int
	for i := 0; i < 260; i++ {
		cmap[i] = min
	}
	for i := 0; i < len(t); i++ {
		cmap[t[i]] = 1
	}
	min = len(s)
	for i := 0; i < len(s); i++ {
		if cmap[s[i]] > -1 {
			cmap[s[i]]++
		}
	}
	for i := 0; i < len(t); i++ {
		cmap[t[i]]--
	}
	for i := 0; i < len(t); i++ {
		if cmap[t[i]] <= 0 {
			return ""
		}
	}
	for mr > ml {
		if cmap[s[mr]] > -1 {
			if cmap[s[mr]] == 1 {
				break
			} else {
				cmap[s[mr]]--
			}
		}
		mr--
	}
	min = mr - ml
	l = ml
	r = mr
	for l <= r {
		if cmap[s[l]] > -1 {
			if cmap[s[l]] == 1 {
				if min > r-l {
					min = r - l
					ml = l
					mr = r
				}
				for r++; r < len(s); r++ {
					if s[r] == s[l] {
						break
					}
					if cmap[s[r]] > -1 {
						cmap[s[r]]++
					}
				}
				l++
				if r == len(s) {
					break
				}
			} else {
				cmap[s[l]]--
				l++
			}
		} else {
			l++
		}
	}

	return s[ml : ml+min+1]
}
func main() {
	//fmt.Println(isNumber("123456"))

	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println("Hello")
}
