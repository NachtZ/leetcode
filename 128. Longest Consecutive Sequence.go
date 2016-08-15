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

func longestConsecutive(nums []int) int {
    m := make(map[int]bool)
    max := 0
    for _,v := range nums{
        m[v] = true
    }
    for _,v := range nums{
        _,ok := m[v-1]
        if ok == false{
            t := v+1 
            for _,ok = m[t];ok;_,ok = m[t]{
                t++
            }
            if max < t-v{
                max = t-v
            }
        }
    }
    return max
}
func main(){
    fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
    fmt.Print("Hello")
}
