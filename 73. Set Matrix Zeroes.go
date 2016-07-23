package main
import "fmt"
func setZeroes(matrix [][]int)  {
    m := len(matrix)
    if m ==0{
        return
    }
    n := len(matrix[0])
    row,col := make([]bool,m),make([]bool,n)
    for i := 0;i<m;i++{
        for j:=0;j<n;j++{
            if matrix[i][j] == 0{
                row[i] = true
                col[j] = true
            }
        }
    }
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            if row[i] || col[j]{
                matrix[i][j] = 0
            }
        }
    }
    
}
func main(){
    //fmt.Println(minDistance("ab","a"))
    fmt.Print("Hello")
}
