package main
import "fmt"
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
var ret []int

func help(root * TreeNode){
    if root == nil{
        return 
    }
    help(root.Left)
    ret = append(ret, root.Val)
    help(root.Right)
}
func inorderTraversal(root *TreeNode) []int {
    ret = []int{}
    help(root)
    return ret
}
func main(){
    fmt.Println(restoreIpAddresses("25525511135"))
    fmt.Print("Hello")
}
