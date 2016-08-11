package main
import "fmt"
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
var ret []*TreeNode
func help(start,end int)[]*TreeNode{
    var res []*TreeNode
    if start > end{
        return []*TreeNode{nil}
    }
    for i:=start;i<=end;i++{
        lefts := help(start,i-1)
        rights := help(i+1,end)
        for j:=0;j<len(lefts);j++{
            for k:=0;k<len(rights);k++{
                root := new(TreeNode)
                root.Val = i
                root.Left = lefts[j]
                root.Right = rights[k]
                res = append(res,root)
            }
        }
    }
    return res
}
func generateTrees(n int) []*TreeNode {
    if n == 0{
        root := new(TreeNode)
        root.Val = 0
        root.Left = nil
        root.Right = nil
        return []*TreeNode{root}
    }
    return help(1,n)
}
func main(){
    fmt.Println(restoreIpAddresses("25525511135"))
    fmt.Print("Hello")
}
