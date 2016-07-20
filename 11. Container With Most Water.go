
func max(a,b int) int{
    if a>b{
        return a
    }
    return b
}
func min(a,b int) int{
    if a>b{
        return b
    }
    return a
}
func maxArea(height []int) int {
    ans :=0
    i,j := 0,len(height)-1
    for i<j{
        ans = max(ans,(j-i)*min(height[i],height[j]))
        if height[i] > height[j]{
            j--
        }else {
            i++
        }
    }
    return ans
}