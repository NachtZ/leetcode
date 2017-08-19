
package main
import "fmt"
import "sort"

var ret [][]int
var now []int 
func help(c []int,sum,target int, idx int){
    if sum == target{
        tmp := make([]int,len(now))
        copy(tmp,now)
        ret = append(ret, tmp)
        return 
    }
    if sum > target{
        return
    }
    for i:=idx;i<len(c);i++{
        now = append(now,c[i])
        help(c,sum+c[i],target,i)
        now = now[:len(now)-1]
    }
}

func combinationSum(candidates []int, target int) [][]int {
    ret = [][]int{}
    now = []int{}
    sort.Ints(candidates)
    help(candidates,0,target,0);
    return ret
}
func main(){
    c := []int{2,3,6,7}
    fmt.Println(combinationSum(c,7))
}
