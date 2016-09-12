func integerReplacement(n int) int {
    if n <= 1{
        return 0
    }
    if n %2 == 0{
        return integerReplacement(n/2)+1
    }else{
        t1,t2 := integerReplacement(n+1),integerReplacement(n-1)
        if t1 <t2{
            return t1 +1
        }else{
            return t2 +1
        }
    }
}