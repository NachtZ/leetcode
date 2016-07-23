package main
import "fmt"
func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    if m == 0{
        return false
    }
    n := len(matrix[0])
    if n == 0{
        return false
    }
    if matrix[m-1][0] == target{
        return true
    }
    for i,j:=0,n-1;i<=m-1 && j >=0;{
        if target == matrix[i][j]{
            return true
        }
        if matrix[i][j] < target{
            i++
        }else{
            j--
        }
    }
    return false
}
func main(){
    fmt.Println(searchMatrix([][]int{{1,1},{2,2}},1))
    fmt.Print("Hello")
}
