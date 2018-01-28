package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func dfs(v []int, ret *string, n, k, l int) bool {
	if len(*ret) == l {
		return true
	}
	*ret += "0"
	for i := 0; i < k; i++ {
		*ret = (*ret)[:len(*ret)-1]
		*ret += string(i + '0')
		t, _ := strconv.Atoi(string((*ret)[len(*ret)-n:]))
		if v[t] == 0 {
			v[t] = 1
			if dfs(v, ret, n, k, l) {
				return true
			}
			v[t] = 0
		}
	}
	*ret = (*ret)[:len(*ret)-1]
	return false
}

func crackSafe(n int, k int) string {
	ret := strings.Repeat("0", n)
	if k == 1 {
		return ret
	}
	N := 10000
	v := make([]int, N)
	ret += "1"
	v[0], v[1] = 1, 1
	l := int(math.Pow(float64(k), float64(n))) + n - 1
	dfs(v, &ret, n, k, l)
	return ret
}
func main() {
	start := time.Now()
	res := crackSafe(2, 2)
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
