package main

import (
	"fmt"

	//"math/rand"
	//"sort"
)

func getParant(e []int, n int) int{
	if (e[n] != n){
		e[n] = getParant(e, e[n]);
	}
	return e[n];
}

func repeatedStringMatch(A string, B string) int {
    for i,j := 0,0;i<len(A);i++{
		for j = 0;j< len(B) && A[(i+j)%len(A)] == B[j];j++{
		}
		if j == len(B){
			if (i+j)%len(A) != 0{
				return (i+j)/len(A) + 1
			}else{
				return (i+j)/len(A)
			}
		}
	}
	return -1
}
func main() {
	fmt.Println(repeatedStringMatch("abcd","cdabcdab"))
}
