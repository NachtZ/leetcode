var m map[int]int

func help(stones []int, idx int, stp int) bool {
	if stp == 0 {
		return false
	}
	if idx == len(stones)-1 {
		return true
	}
	if m[stones[idx]+stp] == 0 {
		return false
	}
	idx = m[stones[idx]+stp]
	return help(stones, idx, stp-1) || help(stones, idx, stp+1) || help(stones, idx, stp)

}
func canCross(stones []int) bool {
	m = make(map[int]int)
	for i := 0; i < len(stones); i++ {
		m[stones[i]] = i
	}
	return help(stones, 0, 1)
}