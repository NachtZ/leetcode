package main
import "fmt"
//import "unicode"
//import "strings"
/***************************************/
//some common data struct and function using in leetcode.
type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
  type ListNode struct {
      Val int
     Next *ListNode
  }

 type Point struct {
      X int
      Y int
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
        if i==len(nums)||nums[i] == -1{
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

func buildList(nums []int) *ListNode{
    if len(nums) == 0{
        return nil
    }
    root := &ListNode{
        Val:nums[0],
    }
    tmp := root
    for i:=1;i<len(nums);i++{
        tmp.Next = &ListNode{
            Val:nums[i],
        }
        tmp = tmp.Next
    }
    return root
}

/**************************************************************/
func maxPoints(points []Point) int {
    if len(points) <=2{
        return len(points)
    }
    
    max :=0
    for i:=0;i<len(points);i++{
        line := 0
        base :=1
        tmax :=0
        val :=0.0
        count := make(map[float64]int)
        for j:=0;j<len(points);j++{
            if i == j{
                continue
            }
            if points[i].X == points[j].X{
                if points[i].Y == points[j].Y{
                    base ++
                    continue
                }
                line ++
                if tmax <line{
                    tmax = line
                }
            }else{
                val = float64(points[i].Y-points[j].Y)/float64(points[i].X - points[j].X)
                if _,ok := count[val];ok == false{
                    count[val] = 0
                }
                count[val]++
                if tmax < count[val]{
                    tmax = count[val]
                }
            }
        }
        tmax = tmax +base
        if max < tmax{
            max = tmax
        }
    }
    return max
}
func main(){
    fmt.Println(maxPoints([]Point{}))
    fmt.Print("Hello")
}
