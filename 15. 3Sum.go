
func threeSum(nums []int) [][]int {
	var ret [][]int
	if len(nums) < 3 {
		return ret
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		target := 0 - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			hit := target - nums[j]
			if hit < nums[j] {
				continue
			}
			for k := j + 1; k < len(nums) && hit >= nums[k]; k++ {
				if hit == nums[k] {
					tmp := []int{nums[i], nums[j], nums[k]}
					ret = append(ret, tmp)
					break
				}
			}
		}
	}
	return ret
}
