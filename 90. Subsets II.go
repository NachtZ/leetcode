package main
import "fmt"
import "sort"
var res [][]int

func dfs(nums []int,start,num int, tmp []int){
    if len(tmp) == num{
        tmp1 := make([]int,len(tmp))
        copy(tmp1,tmp)
        res = append(res,tmp1)
        return 
    }
    for i:=start;i<len(nums);i++{
        if i>start && nums[i] == nums[i-1]{
            continue
        }
        tmp = append(tmp,nums[i])
        dfs(nums,i+1,num,tmp)
        tmp = tmp[:len(tmp)-1]
    }
}

func subsetsWithDup(nums []int) [][]int {
    res = [][]int{}
    sort.Ints(nums)
    for i:=0;i<=len(nums);i++{
        tmp := []int{}
        dfs(nums,0,i,tmp)
    }
    return res
}
func main(){
    fmt.Println(subsetsWithDup([]int{1,2,2}))
    fmt.Print("Hello")
}
