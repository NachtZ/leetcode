//这个主要是将map转为的array。速度60ms。但是感觉还是有点问题。主要是slice的拷贝。
func findSubstring(s string, words []string) []int {
	ret := []int{}
	if len(words) == 0 || len(s) == 0 {
		return ret
	}
	l := len(words[0])
	size := l * len(words)
	if size > len(s) {
		return ret
	}
	idx := 0
	idxmap := make(map[string]int)
	for _, str := range words {
		_, ok := idxmap[str]
		if ok == false {
			idxmap[str] = idx
			idx++
		}
	}
	rawmap := make([]int, len(idxmap))
	for i := 0; i < len(rawmap); i++ {
		rawmap[i] = 0
	}
	for _, str := range words {
		rawmap[idxmap[str]]++
	}

	stridx := make([]int, len(s)-l+1)
	for i := 0; i < len(stridx); i++ {
		str := s[i : i+l]
		idx, ok := idxmap[str]
		if ok == false {
			stridx[i] = -1
		} else {
			stridx[i] = idx
		}
	}
	for i := 0; i < len(stridx); i++ {
		curmap := make([]int, len(rawmap))
		for i := 0; i < len(rawmap); i++ {
			curmap[i] = rawmap[i]
		}
		for j := 0; j < size && i+j < len(stridx); j += l {
			if stridx[i+j] == -1 || curmap[stridx[i+j]] == 0 {
				break
			} else {
				curmap[stridx[i+j]]--
			}
			if j == size-l {
				ret = append(ret, i)
			}
		}
	}
	return ret
}