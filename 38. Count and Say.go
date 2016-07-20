
func help(s string) string{
    res := ""
    char := s[0]
    count := 1
    i:=1
    for i=1;i<len(s);i++{
        if char!=s[i]{
            res += string(byte(count + '0'))
            res += string(char)
            count = 1
            char = s[i]
        }else{
            count ++
        }
    }
    res += string(byte(count + '0'))
    res += string(char)
    return res
}

func countAndSay(n int) string {
    str := "1"
    for i:=1;i<n;i++{
        str = help(str)
    }
    return str
}