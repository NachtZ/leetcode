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
func insertionSortList(head *ListNode) *ListNode {
    preHead := &ListNode{
        Next:head,
    }
    before,now,tail,flag := preHead,head,head,head
    for now != nil{
        flag = now.Next
        before = preHead
        tail.Next = nil
        for ptr := preHead.Next;ptr!=nil;before,ptr = ptr,ptr.Next{
            if now == ptr{
                break
            }
            if now.Val <=ptr.Val{
                before.Next,now.Next = now,ptr
                break
            }
            if ptr == tail{
                tail.Next,tail = now,now
                tail.Next = nil
                break
            }
            
        }
        now = flag
    }
    return preHead.Next
}
func main(){
    t := insertionSortList(buildList([]int{1}))
    for t!=nil{
        fmt.Println(t.Val)
        t = t.Next
    }
    fmt.Print("Hello")
}
