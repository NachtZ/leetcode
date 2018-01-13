func twoSum(nums []int, target int) []int {
	lenNums := len(nums)
	for i := 0; i < lenNums; i++ {
		for j := i + 1; j < lenNums; j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
