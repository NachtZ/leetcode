package main

import (
	"fmt"
	"sort"
	"strconv"
)

func countOfAtoms(formula string) string {
	st, intst, ele, num := []string{}, []int{}, "", 0
	for _, c := range formula {
		if c <= 'Z' && c >= 'A' {
			if ele != "" {
				st = append(st, ele)
				if num > 0 {
					intst = append(intst, num)
					num = 0
				} else {
					intst = append(intst, 1)
				}
			}
			ele = string(c)
		} else if c <= 'z' && c >= 'a' {
			ele = ele + string(c)
		} else if c <= '9' && c >= '0' {
			num = num*10 + int(c-'0')
		} else if c == '(' {
			if ele != "" {
				st = append(st, ele)
				if num > 0 {
					intst = append(intst, num)
					num = 0
				} else {
					intst = append(intst, 1)
				}
			}
			ele = ""
			st = append(st, string(c))
		} else if c == ')' {
			if ele != "" {
				st = append(st, ele)
				if num > 0 {
					intst = append(intst, num)
					num = 0
				} else {
					intst = append(intst, 1)
				}
			}
			ele = ")"
		}

	}

	if ele != "" {
		st = append(st, ele)
		if num > 0 {
			intst = append(intst, num)
			num = 0
		} else {
			intst = append(intst, 1)
		}
	}

	tst, tintst := []string{}, []int{}

	/******************************************************
	explain in Chinese:
	  使用栈来消除括号：
	  	碰到左括号：字符栈压栈，数字栈不动，注意左括号在数字栈中没有对应数字。
	  	碰到右括号：把字符栈所有字符弹栈，直到碰到左括号，数字栈乘以右括号栈对应的数字，然后再把栈压回去。
	  	碰到字符：不动，继续遍历。
	  完成后括号就被消去了，再用map遍历一遍。
	  ******************************************************/
	idx := 0 //push index of stack intst
	for i := range st {
		if st[i] == ")" {
			k, i := intst[idx], len(tst)-1
			for j := len(tintst) - 1; tst[i] != "("; i-- {
				tintst[j] *= k
				j--
			}
			tst = append(tst[:i], tst[i+1:]...)
			//tintst = tintst[:len(tintst)-1]
			idx++
		} else if st[i] == "(" {
			tst = append(tst, st[i])
		} else {
			tst = append(tst, st[i])
			tintst = append(tintst, intst[idx])
			idx++

		}

	}

	m := make(map[string]int)
	for i := 0; i < len(tst); i++ {
		m[tst[i]] += tintst[i]
	}
	sort.Strings(tst)
	ret := ""
	for i := 0; i < len(tst); i++ {
		if i > 0 && tst[i] == tst[i-1] {
			continue
		}
		ret += tst[i]
		if m[tst[i]] > 1 {
			ret += strconv.Itoa(m[tst[i]])
		}
	}
	return ret
}
func main() {
	fmt.Println(countOfAtoms("H2O"))
}
