package main

import (
	"fmt"
	"math"
	"time"
)

//http://blog.csdn.net/thesnowboy_2/article/details/78180542

func minStickers(stickers []string, target string) int {
	n := len(target)
	m := 1 << uint(n)
	dp := make([]int, m)
	for i := 1; i < m; i++ {
		dp[i] = math.MaxInt32
	}
	for i := 0; i < m; i++ {
		if dp[i] == math.MaxInt32 {
			continue
		}
		for _, s := range stickers {
			sup := i
			for _, c := range s {
				for r := 0; r < n; r++ {
					if target[r] == byte(c) && ((uint(sup)>>uint(r))&1) == 0 {
						sup |= 1 << uint(r)
						break
					}
				}
			}
			if dp[i]+1 < dp[sup] {
				dp[sup] = dp[i] + 1
			}
		}
	}
	if dp[m-1] != math.MaxInt32 {
		return dp[m-1]
	}
	return -1
}
func main() {
	start := time.Now()
	res := minStickers([]string{"with", "example", "science"}, "thehat")
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
