package main

import (
	"fmt"

	//"math/rand"
	//"sort"
)
func check(x,y, N int) bool{
	if x <0 || y <0 || x >=N || y >= N{
		return false
	}
	return true
}
func knightProbability(N int, K int, r int, c int) float64 {
	move := [][]int{[]int{2,1},[]int{-2,1},[]int{2,-1},[]int{-2,-1},[]int{1,2},[]int{-1,2},[]int{-1,-2},[]int{1,-2}}
	dp0, dp1 := make([][]float64,N),make([][]float64,N)
	for i:= 0;i<N;i++{
		dp0[i],dp1[i] = make([]float64,N), make([]float64,N)
	}
	remaincnt := float64(0)
	dp0[r][c] = 1.0
	for t:=0;t< K;t++{
		for i:=0;i<N;i++{
			for j:=0;j<N;j++{
				if dp0[i][j] ==0{
					continue
				}
				for s := 0;s<8;s++{
					x, y:= i+move[s][0],j+move[s][1]
					if check(x,y,N){
						dp1[x][y] += dp0[i][j]/8
					}
				}
			}
		}
		dp0,dp1 = dp1,dp0
		for i:=0;i<N;i++{
			for j:=0;j<N;j++{
				dp1[i][j] =0
			}
		}
	}
	for i:=0;i<N;i++{
		for j:=0;j<N;j++{
			remaincnt += dp0[i][j]
		}
	}


	return remaincnt

}

func main() {
	fmt.Println(knightProbability(8,30,6,4))
}
