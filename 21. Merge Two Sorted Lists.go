
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    head := new(ListNode)
    ptr := head
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val{
            ptr.Next = l1
            l1 = l1.Next
        }else{
            ptr.Next = l2
            l2 = l2.Next
        }
        ptr = ptr.Next
    }
    if l1 != nil{
        ptr.Next = l1
    }else{
        ptr.Next = l2
    }
    return head.Next
}