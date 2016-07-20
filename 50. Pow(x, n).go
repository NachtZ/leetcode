

func myPow(x float64, n int) float64 {
    if n== 0{
        return 1.0
    }
    if n<=0{
        if n ==-2147483647{
            return 1.0/(myPow(x,2147483647)*x)
        }else{
            return 1.0/myPow(x,-n)
        }
    }
    ans := 1.0
    for n>0{
        if n& 1 >0{
            ans *= x
        }
        x *=x
        n >>=1
    }
    return ans
}