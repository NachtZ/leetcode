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

func maxProfit(prices []int) int {
    if len(prices) == 0{
        return 0
    }
    f1,f2 := make([]int,len(prices)),make([]int,len(prices))
    minV := prices[0]
    for i:=1;i<len(prices);i++{
        if minV > prices[i]{
            minV = prices[i]
        }
        if f1[i-1] >prices[i]-minV{
            f1[i] = f1[i-1]
        }else{
            f1[i] = prices[i] - minV
        }
    }
        maxV := prices[len(prices)-1]
        f2[len(f2)-1] = 0
        for i := len(prices) -2 ;i>=0;i--{
            if maxV < prices[i]{
                maxV = prices[i]
            }
            if f2[i+1] > maxV - prices[i]{
                f2[i] = f2[i+1]
            }else{
                f2[i] = maxV - prices[i]
        }
    }
            sum := 0
            for i:= 0;i<len(prices);i++{
                if sum < f1[i]+f2[i]{
                    sum = f1[i]+f2[i]
                }
            }
    return sum
}
func main(){
    fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
    fmt.Print("Hello")
}
