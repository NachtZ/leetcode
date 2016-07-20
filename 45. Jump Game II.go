
func jump(nums []int) int {
    ret :=0
    curMax :=0
    curR :=0
    for i:=0;i<len(nums);i++{
        if curR <i{
            ret ++
            curR = curMax
        }
        if curMax < nums[i]+i{
            curMax = nums[i]+i
        }
    }
    return ret
}