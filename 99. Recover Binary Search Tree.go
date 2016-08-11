package main
import "fmt"
type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
var s1,s2,pre *TreeNode
func help(root *TreeNode){
    if root == nil{
        return
    }
    help(root.Left)
    if pre != nil && pre.Val > root.Val{
        if s1 == nil{
            s1 = pre
            s2 = root
        }else{
            s2 = root
            return 
        }
    }
    pre = root
    help(root.Right)
}
func recoverTree(root *TreeNode)  {
    if root == nil{
        return 
    }
    s1,s2,pre = nil,nil,nil
    help(root)
    tmp := s1.Val
    s1.Val = s2.Val
    s2.Val = tmp

}
func main(){
    fmt.Println(isInterleave("aabcc","dbbca","aadbbcbcac"))
    fmt.Print("Hello")
}
