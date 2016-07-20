
func removeDuplicates(nums []int) int {
        if len(nums) <=1{
        return len(nums)
    }
    l :=1
    for i:=1;i<len(nums);i++{
        if nums[i]!=nums[l-1]{
            nums[l] = nums[i]
            l ++
        }
    }
    return l
}