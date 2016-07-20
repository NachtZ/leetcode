//注意在使用全局变量的时候，初始化的时候要记得清空。

var ret []string

func help(btn []string, str, digits string, idx int) {
	if idx == len(digits) {
		ret = append(ret, str)
		return
	}
	tmp := btn[digits[idx]-48]
	for _, t := range tmp {
		help(btn, str+string(t), digits, idx+1)
	}
}

func letterCombinations(digits string) []string {
	btn := []string{" ", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	ret = []string{}
	if len(digits) > 0 {
		help(btn, "", digits, 0)
	}
	return ret
}