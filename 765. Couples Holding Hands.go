package main

import (
	"fmt"
	"time"
)

func minSwapsCouples(row []int) int {
	idx := make(map[int]int)
	for i, v := range row {
		idx[v] = i
	}
	ret := 0
	for i := 0; i < len(row); i += 2 {
		tmp := idx[row[i]+1]
		if row[i]%2 == 1 {
			tmp = idx[row[i]-1]
		}
		if tmp == i+1 {
			continue
		} else {
			row[i+1], row[tmp] = row[tmp], row[i+1]
			idx[row[i+1]], idx[row[tmp]] = i+1, tmp
			ret++
		}
	}
	return ret

}
func main() {
	start := time.Now()
	res := minSwapsCouples([]int{0, 1, 2, 3})
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
