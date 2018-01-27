package main

import (
	"fmt"
	"time"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	m := make(map[rune]int)
	for _, l := range licensePlate {
		if l >= 'A' && l <= 'Z' {
			m[l-'A'+'a']++
		} else if l >= 'a' && l <= 'z' {
			m[l]++
		}
	}
	idx := -1
	for i := range words {
		cnt, flag := make(map[rune]int), true
		for _, w := range words[i] {
			cnt[w]++
		}
		for k, v := range m {
			if cnt[k] == 0 || cnt[k] < v {
				flag = false
				break
			}
		}
		if flag {
			if idx == -1 {
				idx = i
			} else if len(words[idx]) > len(words[i]) {
				idx = i
			}
		}
	}
	return words[idx]
}

func main() {
	start := time.Now()
	res := shortestCompletingWord("Ah71752", []string{"suggest", "letter", "of", "husband", "easy", "education", "drug", "prevent", "writer", "old"})
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
