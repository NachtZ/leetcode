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
func flatten(root *TreeNode)  {
    if root == nil{
        return
    }
    flatten(root.Right)
    flatten(root.Left)
    if root.Left != nil{
        tmp := root.Left
        for tmp != nil && tmp.Right != nil{
            tmp = tmp.Right
        }
        tmp.Right = root.Right
        root.Right = root.Left
        root.Left = nil
    }
}

func main(){
    fmt.Println(pathSum(buildTree([]int{5,4,8,11,-1,13,4,7,2,-1,-1,5,1}),22))
    fmt.Print("Hello")
}
