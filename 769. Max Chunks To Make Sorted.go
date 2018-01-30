package main

import (
	"fmt"
	"sort"
	"time"
)

func maxChunksToSorted(arr []int) int {
	ret, sum1, sum2 := 0, 0, 0
	arr_pre_max := arr[0]
	tarr := make([]int, len(arr))
	copy(tarr, arr)
	sort.Ints(tarr)
	for i := 0; i < len(arr); i++ {
		x, y := arr[i], tarr[i]
		sum1, sum2 = sum1+x, sum2+y
		if x > arr_pre_max {
			arr_pre_max = x
		}
		if arr_pre_max == y && sum1 == sum2 {
			ret++
		}
	}
	return ret
}

func main() {
	start := time.Now()
	res := maxChunksToSorted([]int{5, 4, 3, 2, 1})
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
