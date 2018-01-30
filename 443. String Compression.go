package main

import (
	"fmt"
	"strconv"
	"time"
)

func compress(chars []byte) int {
	if len(chars) < 2 {
		return len(chars)
	}
	pre, ct, size := -1, 1, len(chars)
	c := chars[0]
	for i := 1; i < len(chars); i++ {
		if chars[i] == chars[i-1] {
			ct++
		} else {
			pre++
			chars[pre] = c
			if ct > 1 {
				s := strconv.Itoa(ct)
				for j := 0; j < len(s); j, pre = j+1, pre+1 {
					chars[pre+1] = s[j]
				}

			}
			ct, c = 1, chars[i]
		}
	}
	pre++
	chars[pre] = c
	if ct > 1 {
		s := strconv.Itoa(ct)
		for j := 0; j < len(s); j, pre = j+1, pre+1 {
			chars[pre+1] = s[j]
		}
	}
	for i := pre + 1; i < size; i++ {
		chars = chars[:len(chars)-1]
	}
	fmt.Println(chars)
	return pre + 1
}
func main() {
	start := time.Now()
	res := compress([]byte{'a', 'b', 'b', 'b', 'b'})
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
