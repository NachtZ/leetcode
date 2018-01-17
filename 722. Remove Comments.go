package main

import (
	"fmt"
)

///*
//123*/
func removeComments(source []string) []string {
	var ret []string
	star := false
	v := ""
	for _, s := range source {

		for i := 0; i < len(s); {
			if star {
				if i+1 == len(s) {
					i++
				} else {
					t := string(s[i : i+2])
					if t == "*/" {
						star, i = false, i+2
					} else {
						i++
					}
				}
			} else {
				if i+1 == len(s) {
					v += string(s[i])
					i++
				} else {
					t := string(s[i : i+2])
					if t == "/*" {
						star, i = true, i+2
					} else if t == "//" {
						break
					} else {
						v += string(s[i])
						i++
					}
				}
			}
		}
		if v != "" && star == false {
			ret = append(ret, v)
			v = ""
		}
	}
	return ret
}
func main() {
	fmt.Println(removeComments([]string{"/*Test program */", "int main()", "{ ", "  // variable declaration ", "int a, b, c;", "/* This is a test", "   multiline  ", "   comment for ", "   testing */", "a = b + c;", "}"}))
}
