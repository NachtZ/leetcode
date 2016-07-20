
func reverse(head * ListNode) * ListNode{
    l := new(ListNode)
    tmp := head
    l.Next = head
    for head!=nil{
        tmp = head.Next
        head.Next = l.Next
        l.Next = head
        head = tmp
    }
    return l.Next
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    count :=1
    var tmp,end,before *ListNode
    changed := false
    newHead := head
    tmp = head
    for head!=nil{
        if count == k{
            count  = 1
            end = head.Next
            head.Next = nil
            reverse(tmp)
            if changed == false{
                newHead = head
                changed = true
            }else{
                before.Next.Next = end
                before.Next =head
            }
            before = tmp
            tmp.Next = end
            tmp = end
            head = end
        }else{
            head = head.Next
            count ++
        }

    }
    return newHead
}