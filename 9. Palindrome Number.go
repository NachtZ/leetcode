//注意负数不是回文数

func isPalindrome(x int) bool {
    if x<0{
        return false
    }
    start,end := 0,0
    tmp := x
    ten := 1
    for tmp!=0{
        tmp/=10
        ten *= 10
    }
    ten /=10
    for x!=0{
        start = x/ten
        end = x%10
        if start !=end{
            return false
        }
        x%=ten
        x/=10
        ten/=100
    }
    return true
}