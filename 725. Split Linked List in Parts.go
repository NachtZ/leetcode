package main

import (
	"fmt"
)

func splitListToParts(root *ListNode, k int) []*ListNode {
	size := 0
	for ptr := root; ptr != nil; size, ptr = size+1, ptr.Next {
	}
	ret := []*ListNode{}
	avgLen, left := size/k, size
	for ptr, before := root, root; k > 0; k-- {
		ret = append(ret, ptr)
		for i := 0; i < avgLen && ptr != nil; i, before, ptr = i+1, ptr, ptr.Next {
		}
		if left > avgLen*k {
			left--
			if ptr != nil {
				before, ptr, ptr.Next = ptr, ptr.Next, nil
			} else {
				before, ptr = nil, nil
			}
		} else if before != nil {
			before.Next = nil
		}
		left -= avgLen

	}
	return ret
}
func main() {
	fmt.Println(splitListToParts(buildList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), 3))
}
