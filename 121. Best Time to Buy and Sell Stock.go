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

func maxProfit(prices []int) int {
    dp := make([]int,len(prices))
    dp[0] = 0
    max := 0
    for i:=1;i<len(prices);i++{
        dp[i] = dp[i-1] + prices[i] - prices[i-1]
        if dp[i]<0{
            dp[i] = 0
        }
        if max < dp[i]{
            max = dp[i]
        }
    }
    return max
}
func main(){
    fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
    fmt.Print("Hello")
}
