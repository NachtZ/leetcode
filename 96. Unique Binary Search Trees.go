package main
import "fmt"


func numTrees(n int) int {
    if n <=2{
        return n
    }
    dp := []int{1,1,2}
    for i:=3;i<=n;i++{
        t :=0
        for j:=1;j<=i;j++{
            t += dp[j-1]*dp[i-j]
        }
        dp = append(dp,t)
    }
    return dp[n]
}
func main(){
    fmt.Println(numTrees(5))
    fmt.Print("Hello")
}
