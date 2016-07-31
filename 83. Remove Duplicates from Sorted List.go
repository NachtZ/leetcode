package main
import "fmt"
  type ListNode struct {
      Val int
      Next *ListNode
  }

func deleteDuplicates(head *ListNode) *ListNode {
    var h ListNode
    //h := &tmpelem
    h.Next = head
    for head != nil && head.Next !=nil{
        if head.Val == head.Next.Val{
            head.Next = head.Next.Next
        }else{
            head = head.Next
        }
    }
    return h.Next
}
func main(){
  //  fmt.Println(searchMatrix([][]int{{1,1},{2,2}},1))
    fmt.Print("Hello")
}
