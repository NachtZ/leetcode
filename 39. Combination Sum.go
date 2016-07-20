
package main
import "fmt"
import "sort"

func help(ret [][]int,c []int,now []int,sum,target int, idx int){
    if sum == target{
        //tmp := make([]int,len(now))
        //copy(tmp,now)
        //fmt.Println(&tmp[0],tmp)
        //fmt.Println(&now[0],now)
        ret = append(ret, now)
        fmt.Println(now)
    }
    if sum > target{
        return
    }
    for i:=idx;i<len(c);i++{
        if i> idx && c[i] == c[i-1]{
            continue
        }
        now = append(now,c[i])
        help(ret,c,now,sum+c[i],target,i)
        now = now[:len(now)-1]
    }
    
}

func combinationSum(candidates []int, target int) [][]int {
 //   ret := [][]int{}
    var ret [][]int
    var now []int
//    now := []int{}
    sort.IntsAreSorted(candidates)
    help(ret,candidates,now,0,target,0);
    return ret
}

func main(){
    c := []int{2,3,6,7}
    fmt.Println(combinationSum(c,7))
}
