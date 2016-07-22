package main

import "fmt"

func getPermutation(n int, k int) string {
	size := []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800}

	list := make([]int, n)
	for i := 1; i <= n; i++ {
		list[i-1] = i
	}
	str := ""
	k--
	for i := 0; i < n; i++ {
		total := size[n-i]
		idx := int(k / int(total/(n-i)))
		idx %= n - i
		t := list[idx]
		list = append(list[:idx], list[idx+1:]...)
		str += string('0' + t)

	}
	return str
}
func main() {
	fmt.Println(getPermutation(4, 1))
}
