package main

import (
	"fmt"
	"time"
)

func openLock(deadends []string, target string) int {
	m := make(map[string]bool)
	for _, v := range deadends {
		m[v] = true
	}
	if m["0000"] == true {
		return -1
	}
	if target == "0000" {
		return 0
	}
	v := make(map[string]bool)
	ch := []string{}
	ch = append(ch, "0000")
	for time := 1; len(ch) != 0; time++ {
		for n := len(ch); n > 0; n-- {
			cur := ch[0]
			ch = ch[1:]
			for i := 0; i < 4; i++ {
				for dif := 1; dif <= 9; dif += 8 {
					s := []rune(cur)
					s[i] = rune((int(s[i])-int('0')+dif)%10 + int('0'))
					if string(s) == target {
						return time
					}
					if m[string(s)] != true && v[string(s)] != true {
						ch = append(ch, string(s))
					}
					v[string(s)] = true
				}
			}
		}
	}
	return -1
}
func main() {
	start := time.Now()
	res := openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202")
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
