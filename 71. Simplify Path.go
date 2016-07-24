package main

import "fmt"
import "strings"

func splite(s rune) bool {
	if s == '/' {
		return true
	}
	return false
}

func simplifyPath(path string) string {
	dirs := strings.FieldsFunc(path, splite)
	for i := 0; i < len(dirs); i++ {
		if dirs[i] == "." {
			dirs = append(dirs[:i], dirs[i+1:]...)
			i--
		} else {
			if dirs[i] == ".." {
				if i == 0 {
					dirs = dirs[1:]
					i--
					continue
				}
				dirs = append(dirs[:i-1], dirs[i+1:]...)
				i -= 2
			}
		}
	}
	res := "/"
	for i := 0; i < len(dirs); i++ {
		res += dirs[i]
		if i != len(dirs)-1 {
			res += "/"
		}
	}
	return res
}
func main() {
	//fmt.Println(isNumber("123456"))

	fmt.Println(simplifyPath("/.."))
	fmt.Println("Hello")
}
