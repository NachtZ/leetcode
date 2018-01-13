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

func findRedundantConnection(edges [][]int) []int {
	e :=make([]int,2005)
	for i:=0;i<2005;i++{
		e[i] = i
	}
	for _,edge := range edges{
		p,l := edge[0],edge[1]
		if getParant(e,p) == getParant(e,l){
			return edge
		}else{
			e[getParant(e,p)] = getParant(e,l)
		}
	}
	return nil

}
func main() {
	fmt.Println(findRedundantConnection([][]int{[]int{1,2},[]int{1,3},[]int{2,3}}))
}
