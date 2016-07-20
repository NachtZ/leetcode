
func help(nums []int, left, right, t int) int {
	if left >= right {
		if nums[left] != t {
			return -1
		} else {
			return left
		}
	}
	mid := int((left + right) / 2)
	ret := help(nums, left, mid, t)
	if ret != -1 {
		return ret
	}
	return help(nums, mid+1, right, t)
}
func searchRange(nums []int, target int) []int {
	t := help(nums, 0, len(nums)-1, target)
	if t == -1 {
		return []int{-1, -1}
	}
	left, right := t, t
	for left = t; left >= 0 && nums[left] == target; left-- {
	}
	if left < 0 || nums[left] != target {
		left++
	}
	for right = t; right < len(nums) && nums[right] == target; right++ {

	}
	if right == len(nums) || nums[right] != target {
		right--
	}
	return []int{left, right}
}