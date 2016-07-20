
func removeElement(nums []int, val int) int {
    l :=0
    for i:=0;i<len(nums);i++{
        if nums[i]!=val{
            nums[l] = nums[i]
            l++
        }
    }
    return l
}