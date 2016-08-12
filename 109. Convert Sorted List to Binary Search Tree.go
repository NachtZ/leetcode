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
func help(left,right int) * TreeNode{
    if left >right{
        return nil
    }
    idx :=int(left + (right-left)/2)
    leftNode := help(left,idx-1)
    root :=new(TreeNode)
    root.Val = nums.Val
    root.Left = leftNode
    nums = nums.Next
    root.Right = help(idx +1,right)
    return root
}

func sortedListToBST(head *ListNode) *TreeNode {

    nums=head
    lens := 0
    for head!=nil{
        lens ++
        head = head.Next
    }
    return help(0,lens-1)
}
func main(){
    buildTree([]int{1,2},[]int{1,2})
    fmt.Print("Hello")
}
