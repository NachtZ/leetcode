
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

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) ==0 {
        return nil
    }
    if len(lists) == 1{
        return lists[0]
    }
    ret := lists[0]
    for i :=1;i<len(lists);i++{
        ret = mergeTwoLists(ret,lists[i])
    }
    return ret
}