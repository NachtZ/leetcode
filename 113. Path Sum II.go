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
var path []int
var ret [][]int
func help(root * TreeNode, sum int){
    if root == nil{
        return 
    }
    sum -= root.Val
    path = append(path,root.Val)
    if root.Left == nil && root.Right == nil{
        if sum == 0{
            tmp := make([]int,0)
            tmp = append(tmp,path...)
            ret = append(ret,tmp)
        }
        path = path[:len(path)-1]
        return
    }
    help(root.Left,sum)
    help(root.Right,sum)
    path = path[:len(path)-1]
}

func pathSum(root *TreeNode, sum int) [][]int  {
    path = []int{}
    ret = [][]int{}
    help(root,sum)
    return ret
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
func main(){
    fmt.Println(pathSum(buildTree([]int{5,4,8,11,-1,13,4,7,2,-1,-1,5,1}),22))
    fmt.Print("Hello")
}
