package main

import "fmt"

//96:0+0,97:0+1;98:1+1
func addBinary(a string, b string) string {
	la := len(a)
	lb := len(b)
	if lb > la {
		tmp := a
		a = b
		b = tmp
		la = len(a)
		lb = len(b)
	}
	if la > lb {
		for i := la - lb; i > 0; i-- {
			b = "0" + b
		}
	}
	plus := 0
	res := ""
	for i := la - 1; i != -1; i-- {
		ans := int(a[i]) + int(b[i]) + plus

		if ans >= 98 {
			res = string(ans-'2') + res
			plus = 1
		} else {
			res = string(ans-'0') + res
			plus = 0
		}
	}
	if plus == 1 {
		res = "1" + res
	}
	return res

}
func main() {
	//fmt.Println(isNumber("123456"))
	fmt.Println(addBinary("11", "1"))
	fmt.Println("Hello")
}
