package main
import "fmt"
func generateMatrix(n int) [][]int {
    total := n*n
    idx := 1
    ret := [][]int{}
    
    for i:=0;i<n;i++{
        tmp := make([]int,n)
        ret = append(ret,tmp)
    }
    for i:=0;idx <= total;i++{
        for j:=i;idx <= total && j<n-i;j++{
            ret[i][j] = idx
            idx ++
        }
        for j:=i+1;idx <=total && j<n-i;j++{
            ret[j][n-i-1] = idx
            idx ++
        }
        for j:= n-i-2;idx <=idx&&j>=i;j--{
            ret[n-i-1][j] = idx
            idx ++
        }
        for j:=n-i-2;idx <=idx && j>=i+1;j--{
            ret[j][i] = idx
            idx ++
        }
    }
    return ret
}
func main(){
    fmt.Println(generateMatrix(3))
    fmt.Print("Hello")
}
