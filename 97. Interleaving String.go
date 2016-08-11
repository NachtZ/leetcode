package main
import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
    n1,n2,n3 :=len(s1),len(s2),len(s3)
    if n1+n2 != n3{
        return false
    }
    dp := [][]bool{}
    for i:=0;i<n1+1;i++{
        dp = append(dp,make([]bool,n2+1))
    }
    dp[0][0] = true
    for i:=1;i<=n2;i++{
        dp[0][i] = dp[0][i-1] && s2[i-1] == s3[i-1]
    }
    for i:=1;i<=n1;i++{
        dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
    }
    for i :=1;i<=n1;i++{
        for j:=1;j<=n2;j++{
            dp[i][j] = (dp[i-1][j]&&s1[i-1]==s3[i+j-1])||(dp[i][j-1]&&s2[j-1] == s3[i+j-1])

        }
    }
    return dp[n1][n2]
}
func main(){
    fmt.Println(isInterleave("aabcc","dbbca","aadbbcbcac"))
    fmt.Print("Hello")
}
