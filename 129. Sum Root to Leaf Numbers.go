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

var ret int

func help(root *TreeNode, sum int){
    if root == nil{
        return 
    }
    sum = sum * 10 + root.Val
    if root.Left == nil && root.Right == nil{
        ret += sum
        return
    }
    
    help(root.Left,sum)
    help(root.Right,sum)
}

func sumNumbers(root *TreeNode) int {
    ret = 0
    help(root,0)
    return ret
}
func main(){
    fmt.Println(sumNumbers(buildTree([]int{4,9,0,-1,1})))
    fmt.Print("Hello")
}
