
func romanToInt(s string) int {
    symbol :=[...]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    ret :=0
    value := [...]int{ 1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1 }
    for i,j:=0,0;i<len(symbol)&&j<len(s);{
        if len(symbol[i]) == 2 && j < len(s)-1 {
            if symbol[i][0] == s[j] && symbol[i][1] == s[j+1] {
                j +=2
                ret += value[i]
            }else{
                i++
            }
            continue
        }else if len(symbol[i]) == 1{
            if symbol[i][0] == s[j] {
                j++
                ret += value[i]
            }else{
                i++
            }
        }else{
            i++
        }
    }
    return ret
}