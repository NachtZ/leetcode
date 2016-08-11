package main
import "fmt"

  type ListNode struct {
      Val int
      Next *ListNode
  }

func reverseBetween(head *ListNode, m int, n int) *ListNode {
    var cutHead,cutTail * ListNode
    var newHead,link ListNode
    newHead.Next = head
    link.Next = nil
    tmp := &newHead
    for i :=1;i<m;i++{
        tmp = tmp.Next
    }
    cutHead = tmp
    tmp = tmp.Next
    cutTail = tmp
    for i:=m;i<=n;i++{
        tmp1 := link.Next
        link.Next = tmp
        tmp = tmp.Next
        link.Next.Next = tmp
    }
    if m == 1{
        newHead.Next = link.Next
    }else{
        cutHead.Next = link.Next
    }
    cutTail.Next = tmp
    return newHead.Next
}
func main(){
    fmt.Println(subsetsWithDup([]int{1,2,2}))
    fmt.Print("Hello")
}
