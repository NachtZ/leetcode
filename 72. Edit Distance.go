package main
import "fmt"
func min(a,b int )int{
    if a>b{
        return b
    }
    return a
}
func minDistance(word1 string, word2 string) int {
    m,n := len(word1),len(word2)
    if m ==0 || n ==0{
        if m >n{
            return m
        }
        return n
    }
    m++
    n++
    dp := [][]int{}
    for i:=0;i<m;i++{
         tmp := make([]int,n)
         dp = append(dp,tmp)
    }
    del,ins,rep :=0,0,0
    for i:=0;i<m;i++{
        dp[i][0] = i
    }
    for i:=0;i<n;i++{
        dp[0][i] = i
    }
    m--
    n--
    for i:=1;i<=m;i++{
        for j:=1;j<=n;j++{
            del = dp[i-1][j]+1
            ins = dp[i][j-1]+1
            if word1[i-1] == word2[j-1]{
                rep = dp[i-1][j-1]
            }else{
                rep = dp[i-1][j-1]+1
            }
            dp[i][j] = min(ins,min(del,rep))
        }
    }
    return dp[m][n]

}
func main(){
    fmt.Println(minDistance("ab","a"))
    fmt.Print("Hello")
}
