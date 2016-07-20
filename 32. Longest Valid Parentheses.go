
func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}
	help := make([]bool, len(s))
	stack := []int{}
	for i := len(s) - 1; i >= 0; i-- {
		help[i] = false
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) > 0 {
				help[i] = true
				help[stack[len(stack)-1]] = true
				stack = stack[:len(stack)-1]
			}
		}
	}
	max, l := 0, 0
	for i := 0; i < len(s); i++ {
		if help[i] {
			l++
		} else {
			if max < l {
				max = l
			}
			l = 0
		}
	}
	if max > l {
		return max
	} else {
		return l
	}

}