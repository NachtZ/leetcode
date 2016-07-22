package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	l := 1
	tmp := head
	for tmp.Next != nil {
		l++
		tmp = tmp.Next
	}
	if k >= l {
		k %= l
	}
	if k == 0 {
		return head
	}
	tmp.Next = head
	k = l - k
	for k--; k != 0; k-- {
		head = head.Next
	}
	tmp = head.Next
	head.Next = nil
	return tmp
}
