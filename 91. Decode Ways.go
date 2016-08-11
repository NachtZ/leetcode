package main
import "fmt"

func numDecodings(s string) int {
    
    if len(s) == 0 || s[0] == '0'{
        return 0
    }
    if len(s) == 1{
        return 1
    }
    dp :=make([]int,len(s)+1)
    dp[0] = 1
    if s[1] == '0' && (s[0] == '0'||s[0]>='3'){
            return 0
    }
    if (s[0] == '1' &&s[1]!='0') || (s[0] == '2' && s[1] >'0' && s[1] <='6'){
        dp[1] =2 
    }else{
        dp[1] = 1
    }
    for i:=2;i<len(s);i++{
        if s[i] == '0' && (s[i-1] == '0'||s[i-1]>='3'){
            return 0
        }
        if s[i] == '0'{
            dp[i] = dp[i-2]
        }else{
            if (s[i-1] == '1' &&s[i]!='0') || (s[i-1] == '2' && s[i] >'0' && s[i] <='6'){
                dp[i] = dp[i-1] + dp[i-2]
            }else{
                dp[i] = dp[i-1]
            }
        }
    }
    return dp[len(s)-1]

}
func main(){
    fmt.Println(subsetsWithDup([]int{1,2,2}))
    fmt.Print("Hello")
}
