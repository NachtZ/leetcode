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

func minDepth(root *TreeNode) int {
    stack := []*TreeNode{root}
    depth :=1
    flag := root
    last := root
    tmp := root
    if root == nil{
        return 0
    } 
    
    for len(stack)>0{
        tmp = stack[0]
        if tmp.Left == nil && tmp.Right == nil{
            return depth
        }
        if tmp.Left != nil{
            stack = append(stack,tmp.Left)
            last = tmp.Left
        }
        if tmp.Right != nil{
            stack = append(stack,tmp.Right)
            last = tmp.Right
        }
        if tmp == flag{
            flag = last
            depth ++
        }
        stack = stack[1:]
    }
    return depth
}
func main(){
    buildTree([]int{1,2},[]int{1,2})
    fmt.Print("Hello")
}
