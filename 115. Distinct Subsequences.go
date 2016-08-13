package main
import "fmt"
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
func numDistinct(s string, t string) int {
    if len(s) == 0 || len(t) ==  0{
        return 0
    }
    dp:= make([][]int,len(t)+1)
    for i:=0;i<len(t)+1;i++{
        dp[i] = make([]int,len(s)+1)
    }
    dp[0][0] = 1
    for i:=1;i<=len(s);i++{
        dp[0][i] = 1
    }
    for i:=1;i<=len(t);i++{
        dp[i][0] = 0
    }
    for i:=1;i<=len(t);i++{
        for j:=1;j<=len(s);j++{
            dp[i][j] = dp[i][j-1]
            if s[j-1] == t[i-1]{
                dp[i][j] += dp[i-1][j-1]
            }
        }
    }
    return dp[len(t)][len(s)]
}
func main(){
    fmt.Println(numDistinct("rabbbit","rabbit"))
    fmt.Print("Hello")
}
