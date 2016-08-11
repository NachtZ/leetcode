package main
import "fmt"
type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil{
        return [][]int{}
    }
    list := []*TreeNode{root}
    mark := root
    markNow := root
    flag := true
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
            if flag{
                copy(cpy,level)
            }else{
                for i:=0;i<len(level);i++{
                    cpy[i] = level[len(level)-1-i]                                   
                }
            }
            ret = append(ret,cpy)
            level = []int{}
            flag = !flag
        }
    }
    return ret
}
func main(){
    fmt.Println(isInterleave("aabcc","dbbca","aadbbcbcac"))
    fmt.Print("Hello")
}
