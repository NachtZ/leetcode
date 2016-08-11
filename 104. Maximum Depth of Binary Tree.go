package main
import "fmt"
type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
var deap int

func dfs(root *TreeNode,d int){
    if root == nil{
        if d > deap{
            deap =d
        }
        return
    }
    dfs(root.Left,d+1)
    dfs(root.Right,d+1)
}

func maxDepth(root *TreeNode) int {
    deap = -1
    dfs(root,0)
    return deap    
}
func main(){
    fmt.Println(maxDepth(nil))
    fmt.Print("Hello")
}
