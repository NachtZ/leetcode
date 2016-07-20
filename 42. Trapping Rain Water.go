

func trap(h []int) int {
    if len(h) <=2{
        return 0
    }
    max :=0
    for i:=0;i<len(h);i++{
        if h[max]<h[i]{
            max = i
        }
    }
    left := h[0]
    right := h[len(h)-1]
    cab := 0
    for i:=1;i<max;i++{
        if h[i]>left{
            left = h[i]
        }else{
            cab += left -h[i]
        }
    }
    for i:=len(h)-2;i>max;i--{
        if h[i]>right{
            right = h[i]
        }else{
            cab += right -h[i]
        }
    }
    return cab
}