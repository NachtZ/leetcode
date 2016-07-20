
func intToRoman(num int) string {
    symbol :=[...]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    str := ""
    value := [...]int{ 1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1 }
    for i:=0;num!=0;i++{
        for num >=value[i]{
            num -=value[i]
            str += symbol[i]
        }
    }
    return str
}