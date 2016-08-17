package main
import "fmt"
//import "unicode"
//import "strings"
type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
  type ListNode struct {
      Val int
     Next *ListNode
  }

func buildTree(nums []int)*TreeNode{
    if len(nums) == 0{
        return nil
    }
    root := new(TreeNode)
    root.Val = nums[0]
    ch := make(chan *TreeNode,len(nums))
    ch <- root
    nums =nums[1:]
    for i:=0;i<len(nums);i++{
        tree := <- ch
        if nums[i] == -1{
            tree.Left = nil
        }else{
            tree.Left = &TreeNode{
                Val:nums[i],
            }
            ch <-tree.Left
        }
        i++
        if nums[i] == -1{
            tree.Right = nil
        }else{
            tree.Right = &TreeNode{
                Val:nums[i],
            }
            ch <-tree.Right
        }
    }
    return root
}

func isP(s string,i,j int) bool{
    for ;i<j;i,j = i+1,j-1{
        if s[i] != s[j]{
            return false
        }
    }
    return true
}

func minCut(s string) int {
    dp := make([][]int,len(s))
    for i:=0;i<len(s);i++{
        dp[i] = make([]int,len(s))
        for j:=0;j<len(s);j++{
            dp[i][j] = j-i
        }
    }
    for i:=0;i<len(s);i++{
        for left,right := i,i;left >=0&&right <len(s)&& s[left]==s[right];left,right =left-1,right+1{
            dp[left][right] = 0
        }
        for left,right :=i,i+1;left >=0&&right <len(s)&& s[left]==s[right];left,right =left-1,right+1{
            dp[left][right] =  0
        }
    }
    if dp[0][len(s)-1] ==0{
        return 0
    }
    for i:=1;i<len(s);i++{
        if dp[0][i] == 0{
            continue
        }
        dp[0][i] = dp[0][i-1] +1
        for j:= 0;j<i;j++{
            if dp[0][i] > dp[0][j] + dp[j+1][i]+1{
                dp[0][i] = dp[0][j] +dp[j+1][i]+1
            }
        }
    }
    return dp[0][len(s)-1]
}

func main(){
    fmt.Println(minCut("aab"))
    fmt.Print("Hello")
}
