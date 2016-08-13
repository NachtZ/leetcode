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
var nums *ListNode
func help(root * TreeNode, res *bool)int{
    if *res == false{
        return 0
    }
    if root == nil{
        return 0
    }
    left := help(root.Left,res)
    right := help(root.Right,res)
    if left - right >1 || right - left >1{
        *res = false;
        return 0;
    }
    if left> right{
        return left +1
    }
    return right +1
    

}
func isBalanced(root *TreeNode) bool {
    res := true
    help(root,&res)
    return res
}
func main(){
    buildTree([]int{1,2},[]int{1,2})
    fmt.Print("Hello")
}
