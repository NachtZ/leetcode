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

func getRow(rowIndex int) []int{
    if rowIndex == 0{
        return []int{}
    }
    ret := []int{1}
    if rowIndex == 1{
        return ret
    }
    for i:=1;i<rowIndex;i++{
        ret = append([]int{1},ret...)
        for j:=1;j<i;j++{
            ret[j] = ret[j] + ret[j+1]
        }
    }
    return ret
}
func main(){
    fmt.Println(getRow(5))
    fmt.Print("Hello")
}
