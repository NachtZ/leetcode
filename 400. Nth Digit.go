func findNthDigit(n int) int {
	i := 0
	mark, l := 1, 0
	total := 0
	for i = 1; total < n; i++ {
		if i >= mark {
			total += l + 1
			l++
			mark *= 10
		} else {
			total += l
		}
	}
	i--
	total -= l
	t := strconv.Itoa(i)
	return int(t[n-total-1] - '0')
}