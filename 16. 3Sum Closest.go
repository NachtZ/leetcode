
func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -1 * a
	}

}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	min := nums[0] + nums[1] + nums[2]
	t := abs(min - target)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				res := nums[i] + nums[j] + nums[k]
				if abs(res-target) < t {
					t = abs(res - target)
					min = res
				}
				if res > target {
					break
				}
			}
		}
	}
	return min
}