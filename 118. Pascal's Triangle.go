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

func generate(numRows int) [][]int {
    if numRows == 0{
        return [][]int{}
    }
    ret := [][]int{[]int{1}}
    if numRows == 1{
        return ret
    }
    for i:=1;i<numRows;i++{
        tmp := make([]int,i+1)
        tmp[0] = 1
        for j:=1;j<i;j++{
            tmp[j] = ret[i-1][j-1] + ret[i-1][j]
        }
        tmp[i] = 1
        ret = append(ret,tmp)
    }
    return ret
}
func main(){
    fmt.Println(generate(5))
    fmt.Print("Hello")
}
