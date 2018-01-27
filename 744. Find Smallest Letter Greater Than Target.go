package main

import (
	"fmt"
	"time"
)

func nextGreatestLetter(letters []byte, target byte) byte {
	m := [26]bool{}
	for _, t := range letters {
		m[t-'a'] = true
	}
	for i := 1; i <= 26; i++ {
		if m[(i+int(target)-'a')%26] == true {
			return byte((i+int(target)-'a')%26 + 'a')
		}
	}
	return 'a'
}
func main() {
	start := time.Now()
	res := nextGreatestLetter([]byte{'c', 'f', 'a'}, 'a')
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
