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
func candy(ratings []int) int {
    candies := make([]int,len(ratings))
    for i:=0;i<len(candies);i++{
        candies[i] = 1
    }
    for i:=1;i<len(candies);i++{
        if ratings[i] > ratings[i-1]{
            candies[i] = candies[i-1]+1
        }
    }
    for i:= len(candies) -2 ;i>=0;i--{
        if ratings[i] >ratings[i+1] && candies[i] <= candies[i+1]{
            candies[i] = candies[i+1] + 1
        }
    }
    res := 0
    for i:=0;i<len(candies);i++{
        res += candies[i]
    }
    return res
}

func main(){
    fmt.Println(candy([]int{1,2,3}))
    fmt.Print("Hello")
}
