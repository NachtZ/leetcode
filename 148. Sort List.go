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
func sortList(head *ListNode) *ListNode {
    if head == nil||head.Next == nil{
        return head
    }
    fast,slow,before := head,head,head
    for fast!=nil&&fast.Next != nil{
        fast = fast.Next.Next
        before = slow
        slow = slow.Next
    }
    if fast != nil{
        before = slow
        slow = slow.Next
    }
    before.Next = nil
    list1 := sortList(head)
    list2 := sortList(slow)
    var newHead ListNode
    before = & newHead
    for list1!= nil && list2!= nil{
        if list1.Val <= list2.Val{
            before.Next,before = list1,list1
            list1 = list1.Next
        }else{
            before.Next,before = list2,list2
            list2 = list2.Next
        }
    }
    if list1!=nil{
        before.Next = list1
    }else{
        before.Next = list2
    }
    return newHead.Next
}
func main(){
    t := sortList(buildList([]int{2,1,5,4,3}))
    for t!=nil{
        fmt.Println(t.Val)
        t = t.Next
    }
    fmt.Print("Hello")
}
