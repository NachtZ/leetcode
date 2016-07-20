
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
    for i:=idx+1;i<len(c);i++{
        if i> idx+1 && c[i] == c[i-1]{
            continue
        }
        if c[i] > target{
            break
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
    sort.Ints(candidates)
    for idx,val := range candidates {
        if idx >0 && val == candidates[idx-1]{
            continue
        }
        //if val > target{
        //    break
        //}
        now = append([]int{},val)
        help(ret,candidates,now,val,target,idx);
    }
    return ret
}

func main(){
    c := []int{10, 1, 2, 7, 6, 1, 5}
    fmt.Println(combinationSum(c,8))
}
