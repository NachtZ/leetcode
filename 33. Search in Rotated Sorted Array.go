
//8ms

func search1(nums []int, target int) int {
	for idx, val := range nums {
		if val == target {
			return idx
		}
	}
	return -1
}
//8ms
func help(nums []int, left, right, target int) int {
	if left >= right {
		if nums[left] == target {
			return left
		} else {
			return -1
		}
	}
	if nums[left] < nums[right] {
		if nums[left] > target || nums[right] < target {
			return -1
		}
		mid := int((left + right) / 2)
		ret := help(nums, left, mid, target)
		if ret != -1 {
			return ret
		}
		return help(nums, mid+1, right, target)
	} else {
		mid := int((left + right) / 2)
		ret := help(nums, left, mid, target)
		if ret != -1 {
			return ret
		}
		return help(nums, mid+1, right, target)
	}
	return -1
}

func search2(nums []int, target int) int {
	return help(nums, 0, len(nums)-1, target)
}