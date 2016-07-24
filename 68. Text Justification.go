package main

import "fmt"

func fullJustify(words []string, maxWidth int) []string {
	space := ""
	for i := 0; i < maxWidth; i++ {
		space += " "
	}
	if len(words) == 0 {
		return []string{space}
	}

	ret := []string{}
	row := []string{}
	row = append(row, words[0])
	line := len(words[0])
	size := len(words[0])
	for i := 1; i < len(words); i++ {
		size = len(words[i])
		if line+size+1 > maxWidth {

			line -= len(row) - 1
			numspace := maxWidth - line
			perspace := 0
			more := 0
			if len(row) > 1 {
				perspace = int(numspace / (len(row) - 1))
				more = numspace - perspace*(len(row)-1)
			} else {
				perspace = numspace
			}
			str := ""
			for j := 0; j < len(row); j++ {
				str += row[j]
				if j == 0 || j != len(row)-1 {
					if j < more {
						str += space[:perspace+1]
					} else {
						str += space[:perspace]
					}
				}
			}
			ret = append(ret, str)
			row = []string{words[i]}
			line = len(words[i])
		} else {
			row = append(row, words[i])
			line += size + 1
		}
	}
	str := ""
	for j := 0; j < len(row); j++ {
		str += row[j]
		if j != len(row)-1 {
			str += " "
		}
	}
	str += space[:maxWidth-len(str)]
	ret = append(ret, str)
	return ret
}
func main() {
	//fmt.Println(isNumber("123456"))
	words := []string{"What", "must", "be", "shall", "be."}

	fmt.Println(fullJustify(words, 12))
	fmt.Println("Hello")
}
