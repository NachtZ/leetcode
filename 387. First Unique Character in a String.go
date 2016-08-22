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
func firstUniqChar(s string) int {
    mp := make([]int,26)
    for _,c := range s{
        mp[int(c-'a')]++
    }
    for i,c := range s{
        if mp[int(c-'a')] == 1{
            return i
        }
    }
    return -1

}
func main(){
    fmt.Println(firstUniqChar("loveleetcode"))
    fmt.Print("Hello")
}
