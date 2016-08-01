package main
import "fmt"

func grayCode(n int) []int {
    ret := []int{0,1}
    if n < 2{
        return ret[:n+1]
    }
    base :=2
    n --
    for n>0{
        for i:=len(ret)-1;i>=0;i--{
            ret = append(ret,ret[i]+ base)
        }
        base *=2
        n--
    }
    return ret
}
func main(){
    fmt.Println(grayCode(10))
    fmt.Print("Hello")
}
