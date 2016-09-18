var b = []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512}
var ret []string

func readTime(num int) {
    str := ""
    t := num >> 6
    if t >= 12 {
        return
    }
    str += strconv.Itoa(t) + ":"
    t = num & 63
    if t >= 60 {
        return
    }
    if t < 10 {
        str += "0" + strconv.Itoa(t)
    } else {
        str += strconv.Itoa(t)
    }
    ret = append(ret, str)
}
func help(now, num, idx int) {
    if num == 0 {
        readTime(now)
        return
    }
    if idx >= 10 {
        return
    }
    help(now, num, idx+1)
    help(now|b[idx], num-1, idx+1)
}
func readBinaryWatch(num int) []string {
    ret = []string{}
    if num == 0 {
        return []string{"0:00"}
    }
    for i := 0; i < len(b); i++ {
        help(b[i], num-1, i+1)
    }
    return ret
}