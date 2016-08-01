package main
import "fmt"
  type ListNode struct {
      Val int
      Next *ListNode
  }
 
func isScramble(s1 string, s2 string) bool {
    a := len(s1) +1
    dp :=[][][]bool{}
    for i:= 0;i<a;i++{
        tmp := make([][]bool,0)
        
        for j:=0;j<a;j++{
            tmp1 := make([]bool,a)
            tmp = append(tmp,tmp1)
        }
        dp = append(dp,tmp)
    }
    for i:=1;i<a;i++{
        for j:=1;j<a;j++{
            dp[1][i][j] = s2[j-1] == s1[i-1]
        }
    }
    for k:=2;k<a;k++{
        for t:=1;t<k;t++{
            for i:=1;i<a-k+1;i++{
                for j:=1;j<a-k+1;j++{
                    if (dp[t][i][j] && dp[k-t][i+t][j+t])||(dp[t][i][j+k-t]&&dp[k-t][i+t][j]){
                        dp[k][i][j] = true
                    }
                }
            }
        }
    }
    return dp[a-1][1][1]
}
func main(){

    fmt.Println(isScramble("a","a"))
    fmt.Print("Hello")
}
