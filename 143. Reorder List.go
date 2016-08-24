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

func reorderList(head *ListNode)  {
    if head == nil||head.Next == nil{
        return
    }
    fast, slow := head,head
    before := head
    before = nil
    for fast!=nil && fast.Next !=nil{
        before = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    if fast == nil{
        before.Next = nil
    }else{
        slow = slow.Next
        before.Next.Next = nil
    }
    before = nil
    for slow != nil{
        slow,slow.Next,before = slow.Next,before,slow
    }
    for head!=nil && before!=nil{
        head,head.Next,before,before.Next = head.Next,before,before.Next,head.Next
    }
    return 

}
func main(){
    t := buildList([]int{1,2,3})
    reorderList(t)
    for t!= nil{
        fmt.Println(t.Val)
        t = t.Next
    }
    fmt.Print("Hello")
}
