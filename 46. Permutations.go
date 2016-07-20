
var ret [][]int

func help(nums []int,idx int){
    if idx >=len(nums){
        return
    }
    help(nums,idx+1)
    for i:=0;i<idx;i++{
        tmp := nums[i]
        nums[i] = nums[idx]
        nums[idx] = tmp
        tmp1 := make([]int,len(nums))
        copy(tmp1,nums)
        ret = append(ret,tmp1)
        help(nums,idx+1)
        tmp = nums[i]
        nums[i] = nums[idx]
        nums[idx] = tmp       
    }
}
func permute(nums []int) [][]int {
    ret = [][]int{}
    tmp := make([]int,len(nums))
    copy(tmp,nums)
    ret = append(ret,tmp)
    help(nums,1)
    return ret
}