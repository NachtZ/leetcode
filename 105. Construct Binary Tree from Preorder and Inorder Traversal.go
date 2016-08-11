package main
import "fmt"
type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0{
        return nil
    }

    root := &TreeNode{
        Val:preorder[0],
        Left:nil,
        Right:nil,
    }
    if len(preorder) == 1{
        return root
    }
    idx := 0
    for idx = 0;preorder[0]!=inorder[idx];idx++{
    }
    root.Left = buildTree(preorder[1:idx+1],inorder[:idx])
    root.Right = buildTree(preorder[idx+1:],inorder[idx+1:])
    return root

}
func main(){
    buildTree([]int{1,2},[]int{1,2})
    fmt.Print("Hello")
}
