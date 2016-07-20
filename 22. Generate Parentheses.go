
func help(ret *[]string, left, right int, str string) {
	if left == 0 && right == 0 {
		*ret = append(*ret, str)
	}
	if left > 0 {
		help(ret, left-1, right, str+"(")
	}
	if right > 0 && left < right {
		help(ret, left, right-1, str+")")
	}
}

func generateParenthesis(n int) []string {
	ret := []string{}
	ptr := &ret
	help(ptr, n, n, "")
	return ret
}