package main

import "fmt"
import "strconv"

func monotoneIncreasingDigits(N int) int {
	str := []rune(strconv.Itoa(N))
	m := len(str)
	for i := len(str) - 1; i > 0; i-- {
		if str[i] < str[i-1] {
			m = i
			str[i-1] = rune(str[i-1] - 1)
		}
	}
	for i := m; i < len(str); i++ {
		str[i] = '9'
	}
	ret, _ := strconv.Atoi(string(str))
	return ret
}
func main() {
	fmt.Println(monotoneIncreasingDigits(9))
}
