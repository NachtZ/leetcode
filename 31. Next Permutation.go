

func reverse(nums []int) {
	l := len(nums)
	tmp := 0
	for i := 0; i < l && i < l-i; i++ {
		tmp = nums[i]
		nums[i] = nums[l-i-1]
		nums[l-i-1] = tmp
	}
}

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	var i, k int
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i < 0 {
		reverse(nums)
	}
	for k = len(nums) - 1; i >= 0 && k > i; k-- {
		if nums[i] < nums[k] {
			break
		}
	}
	if i >= 0 {
		tmp := nums[i]
		nums[i] = nums[k]
		nums[k] = tmp
		s := nums[i+1 : len(nums)]
		reverse(s)
		return
	}
}