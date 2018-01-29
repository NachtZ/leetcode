package main

import (
	"fmt"
	"sort"
	"time"
)

func makeLargestSpecial(S string) string {
	if len(S) == 0 {
		return S
	}
	anchor, bal, mountains := 0, 0, []string{}
	for i := 0; i < len(S); i++ {
		if S[i] == '1' {
			bal += 1
		} else {
			bal += -1
		}
		if bal == 0 {
			mountains = append(mountains, "1"+makeLargestSpecial(string(S[anchor+1:1+i]))+"0")
			anchor = i + 1
		}
	}
	sort.Strings(mountains)
	ret := ""
	for i := len(mountains) - 1; i >= 0; i-- {
		ret += mountains[i]
	}
	return ret
}
func main() {
	start := time.Now()
	res := makeLargestSpecial("11011000")
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
