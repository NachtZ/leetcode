
func multiply(num1 string, num2 string) string {
	n1 := make([]int, len(num1))
	n2 := make([]int, len(num2))
	res := make([]int, len(num1)+len(num2))
	for idx, ch := range num1 {
		n1[idx] = int(ch - '0')
	}
	for idx, ch := range num2 {
		n2[idx] = int(ch - '0')
	}
	for i := 0; i < len(num1); i++ {
		for j := 0; j < len(num2); j++ {
			res[i+j+1] += n1[i] * n2[j]
		}
	}
	ss := ""
	for k := len(num1) + len(num2) - 1; k >= 0; k-- {
		if k > 0 {
			res[k-1] += res[k] / 10
		}
		res[k] %= 10
		ss = string(res[k]+'0') + ss
	}
	for len(ss) > 1 && ss[0] == '0' {
		ss = ss[1:]
	}
	return ss
}
```