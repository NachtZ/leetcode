
func swapPairs(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    tmp := head
    list := new(ListNode)
    list.Next = head
    head = list
    for head.Next!=nil && head.Next.Next!=nil{
        tmp = head.Next
        head.Next = tmp.Next
        tmp.Next = head.Next.Next
        head.Next.Next=tmp
        head=head.Next.Next
    }
    return list.Next
}