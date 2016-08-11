package main
import "fmt"
type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
        if root == nil{
        return [][]int{}
    }
    list := []*TreeNode{root}
    mark := root
    markNow := root
    level := []int{}
    ret := [][]int{}
    tmp := root
    for len(list)>0{
        if list[0] == nil{
            list = list[1:]
            continue
        }
        tmp = list[0]
        list = list[1:]
        level = append(level,tmp.Val)
        if tmp.Left != nil{
            markNow = tmp.Left
            list = append(list,tmp.Left)
        }
        if tmp.Right != nil{
            markNow = tmp.Right
            list = append(list,tmp.Right)
        }
        if mark == tmp{
            mark = markNow
            cpy := make([]int,len(level))
            copy(cpy,level)
            ret = append(ret,cpy)
            level = []int{}
        }
    }
    for i,j:= 0,len(ret)-1;i<j;{
        ret[i],ret[j] = ret[j],ret[i]
        i++
        j--
    }
    return ret
}
func main(){
    buildTree([]int{1,2},[]int{1,2})
    fmt.Print("Hello")
}
