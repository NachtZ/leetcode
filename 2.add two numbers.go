/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		l1 = &ListNode{}
	}
	if l2 == nil {
		l2 = &ListNode{}
	}

	result := &ListNode{}
	result.Val = l1.Val + l2.Val

	if result.Val > 9 {
		result.Val -= 10

		if l2.Next == nil {
			l2.Next = &ListNode{}
		}
		l2.Next.Val += 1
	}

	if l2.Next != nil || l1.Next != nil {
		result.Next = addTwoNumbers(l1.Next, l2.Next)
	}
	return result
}
