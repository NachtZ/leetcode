
func firstMissingPositive(nums []int) int {
    rec := make([]int,len(nums)+2)
    l := len(nums)
    for _,val := range nums{
        if val > l || val <=0{
            continue
        }
        rec[val] = 1
    }
    for i :=1;i<=l;i++{
        if rec[i] == 0{
            return i
        }
    }
    return len(nums)+1
}