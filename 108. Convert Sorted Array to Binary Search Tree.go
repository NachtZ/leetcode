package main
import "fmt"
type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

func help(nums []int, left,right int) * TreeNode{
    if left >right{
        return nil
    }
    idx :=int((left+right)/2)
    root := &TreeNode{
        Val:nums[idx],
        Left:nil,
        Right:nil,
    }
    root.Left = help(nums,left,idx-1)
    root.Right = help(nums,idx +1,right)
    return root
}

func sortedArrayToBST(nums []int) *TreeNode {
    return help(nums,0,len(nums)-1)
}
func main(){
    buildTree([]int{1,2},[]int{1,2})
    fmt.Print("Hello")
}
