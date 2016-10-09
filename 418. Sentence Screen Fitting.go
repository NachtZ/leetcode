func wordsTyping(sentence []string, rows int, cols int) int {
	l := make([]int, len(sentence))
	for i := 0; i < len(l); i++ {
		l[i] = len(sentence[i])
	}
	rep := 0
	idx := 0
	left := cols
	//s := 0
	//repeat_start := 0
	for r := 0; r < rows; {

		if left == cols {
			if l[idx] > cols {
				return 0
			}
			left -= l[idx]
			idx++
		} else {
			if left < l[idx]+1 {
				left = cols
				//idx ++
				r++

			} else {
				if l[idx] > cols {
					return 0
				}
				left -= l[idx] + 1
				idx++
			}
		}
		if idx == len(l) {
			rep++
			idx = 0
			if left < l[0]+1 { //a new row
				r++
				tmp := rows / r
				if tmp*r == rows {
					return tmp * rep
				} else {
					rep *= tmp
					r = tmp * r
					r--
				}
			}
		}
	}
	return rep

}