package main
import "fmt"
//import "unicode"
//import "strings"
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
        if i==len(nums)||nums[i] == -1{
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

func buildList(nums []int) *ListNode{
    if len(nums) == 0{
        return nil
    }
    root := &ListNode{
        Val:nums[0],
    }
    tmp := root
    for i:=1;i<len(nums);i++{
        tmp.Next = &ListNode{
            Val:nums[i],
        }
        tmp = tmp.Next
    }
    return root
}

var ret []int

func help(root *TreeNode){
    if root == nil{
        return
    }
    ret = append(ret,root.Val)
    help(root.Left)
    help(root.Right)
}

func preorderTraversal(root *TreeNode) []int {
    ret = []int{}
    help(root)
    return ret
}
func main(){
    fmt.Println(preorderTraversal(buildTree([]int{1,-1,2,3})))
    fmt.Print("Hello")
}
