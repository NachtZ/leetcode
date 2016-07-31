package main
import "fmt"
  type ListNode struct {
      Val int
      Next *ListNode
  }

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil{
        return head
    }
    var h ListNode
    h.Next = head
    now := &h
    num := head.Val
    for head != nil && head.Next!=nil && now.Next !=nil{
        if head.Next.Val == num{
            for head.Next != nil && head.Next.Val == num{
                head = head.Next
            }
            now.Next = head.Next
            if head.Next == nil{
                break
            }
            now.Next = head.Next
            num = head.Next.Val
            head = head.Next
        }else{
            now.Next = head
            now = now.Next
            num = head.Next.Val
            head = head.Next
        }
    }
    return h.Next
}
func main(){
  //  fmt.Println(searchMatrix([][]int{{1,1},{2,2}},1))
    fmt.Print("Hello")
}
