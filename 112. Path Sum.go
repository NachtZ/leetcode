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



func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil{
        return false
    }
    if root.Left == nil && root.Right == nil{
        return root.Val == sum
    }
    sum -= root.Val
    flag := hasPathSum(root.Left,sum)
    if flag == false {
        flag = hasPathSum(root.Right,sum)
    }
    return flag
}
func main(){
    buildTree([]int{1,2},[]int{1,2})
    fmt.Print("Hello")
}
