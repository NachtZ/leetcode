package main
import "fmt"
type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

 var pre * TreeNode
func help(root *TreeNode) bool{
    if root == nil{
        return true
    }
    if help(root.Left) == false{
        return false
    }
    if pre != nil && pre.Val  >= root.Val{
        return false
    }
    pre = root
    return help(root.Right)
}

func isValidBST(root *TreeNode) bool {
    pre = nil
    return help(root)
}
func main(){
    fmt.Println(isInterleave("aabcc","dbbca","aadbbcbcac"))
    fmt.Print("Hello")
}
