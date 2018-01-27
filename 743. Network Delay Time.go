package main

import (
	"fmt"
	"math"
	"time"
)

func networkDelayTime(times [][]int, N int, K int) int {
	dist := make([]int, N+1)
	for i := 0; i < len(dist); i++ {
		dist[i] = math.MaxInt32
	}
	dist[K] = 0
	for i := 0; i < N; i++ {
		for _, t := range times {
			u, v, w := t[0], t[1], t[2]
			if dist[u] != math.MaxInt32 && dist[v] > dist[u]+w {
				dist[v] = dist[u] + w
			}
		}
	}
	mx := 0
	for i := 1; i <= N; i++ {
		if mx < dist[i] {
			mx = dist[i]
		}
	}
	if mx == math.MaxInt32 {
		return -1
	}
	return mx
}
func main() {
	start := time.Now()
	res := networkDelayTime([][]int{[]int{2, 1, 1}, []int{2, 3, 1}, []int{3, 4, 1}}, 4, 2)
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
