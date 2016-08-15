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

func buildTree(nums []int)*TreeNode{
    if len(nums) == 0{
        return nil
    }
    root := new(TreeNode)
    root.Val = nums[0]
    ch := make(chan *TreeNode,len(nums))
    ch <- root
    nums =nums[1:]
    for i:=0;i<len(nums);i++{
        tree := <- ch
        if nums[i] == -1{
            tree.Left = nil
        }else{
            tree.Left = &TreeNode{
                Val:nums[i],
            }
            ch <-tree.Left
        }
        i++
        if nums[i] == -1{
            tree.Right = nil
        }else{
            tree.Right = &TreeNode{
                Val:nums[i],
            }
            ch <-tree.Right
        }
    }
    return root
}

var max int

func help(root * TreeNode){
    if root == nil {
        return
    }
    help(root.Left)
    help(root.Right)
    left,right,max1 := root.Val,root.Val,root.Val
    if root.Left != nil &&root.Left.Val >0{
        left += root.Left.Val
        max1 += root.Left.Val
    }
    if root.Right != nil &&root.Right.Val >0{
        right += root.Right.Val
        max1 += root.Right.Val
    }
    if left > right{
        root.Val = left
    }else{
        root.Val = right
    }
    if max1 > max{
        max = max1
    }
}

func maxPathSum(root *TreeNode) int {
    if root == nil{
        return 0
    }
    max = root.Val
    help(root)
    return max
}
func main(){
    fmt.Println(maxPathSum(buildTree([]int{1,2,3})))
    fmt.Print("Hello")
}
