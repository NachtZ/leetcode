package main
import "fmt"
type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(postorder) == 0{
        return nil
    }
    r := len(postorder) -1
    root := &TreeNode{
        Val:postorder[r],
        Left:nil,
        Right:nil,
    }
    if len(postorder) == 1{
        return root
    }
    idx := 0
    for idx = 0;postorder[r]!=inorder[idx];idx++{
    }
    root.Left = buildTree(inorder[:idx],postorder[0:idx])
    root.Right = buildTree(inorder[idx+1:],postorder[idx:r])
    return root

}
func main(){
    buildTree([]int{1,2},[]int{1,2})
    fmt.Print("Hello")
}
