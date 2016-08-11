package main
import "fmt"
type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
func help(left,right * TreeNode) bool{
    if left == nil && right == nil{
        return true
    }
    if left == nil || right == nil{
        return false
    }
    if left.Val != right.Val {
        return false
    }
    if help(left.Left,right.Right) == false{
        return false
    }
    return help(left.Right,right.Left)
}
func isSymmetric(root *TreeNode) bool {
    if root == nil{
        return true
    }
    return help(root.Left,root.Right)
}
func main(){
    fmt.Println(isInterleave("aabcc","dbbca","aadbbcbcac"))
    fmt.Print("Hello")
}
