package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
	//"math/rand"
	//"sort"
	//"regexp"
)

/***************************************/
//some common data struct and function using in leetcode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type ListNode struct {
	Val  int
	Next *ListNode
}

type Point struct {
	X int
	Y int
}

type Interval struct {
	Start int
	End   int
}

func useLib() {
	fmt.Println(strconv.Itoa(1))
	fmt.Println(strings.Compare("1", "2"))
	fmt.Println(math.Abs(1.0))
	fmt.Println(unicode.IsDigit('1'))
}

func buildTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = nums[0]
	ch := make(chan *TreeNode, len(nums))
	ch <- root
	nums = nums[1:]
	for i := 0; i < len(nums); i++ {
		tree := <-ch
		if nums[i] == -1 {
			tree.Left = nil
		} else {
			tree.Left = &TreeNode{
				Val: nums[i],
			}
			ch <- tree.Left
		}
		i++
		if i == len(nums) || nums[i] == -1 {
			tree.Right = nil
		} else {
			tree.Right = &TreeNode{
				Val: nums[i],
			}
			ch <- tree.Right
		}
	}
	return root
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	root := &ListNode{
		Val: nums[0],
	}
	tmp := root
	for i := 1; i < len(nums); i++ {
		tmp.Next = &ListNode{
			Val: nums[i],
		}
		tmp = tmp.Next
	}
	return root

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/**************************************************************/
/*
type IntSlice []int

func (s IntSlice)Less(i,j int)bool{
	return s[i]<s[j]
}
func (s IntSlice)Swap(i,j int){
	s[i],s[j] = s[j],s[i]
}
func (s IntSlice)Len()int{
	return len(s)
}
*/
/**************************************************************/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	len1, len2 := 0, 0
	for tmp := l1; tmp != nil; tmp = tmp.Next {
		len1++
	}
	for tmp := l2; tmp != nil; tmp = tmp.Next {
		len2++
	}
	if len1 < len2 {
		l1, l2, len1, len2 = l2, l1, len2, len1
	}
	s1, s2 := make([]*ListNode, 0, len1), make([]*ListNode, 0, len2)
	for tmp := l1; tmp != nil; tmp = tmp.Next {
		s1 = append(s1, tmp)
	}
	for tmp := l2; tmp != nil; tmp = tmp.Next {
		s2 = append(s2, tmp)
	}

	var head ListNode
	add, i, t := 0, 0, 0
	for i = 1; i <= len1; i++ {
		if i <= len2 {
			t = s1[len1-i].Val + s2[len2-i].Val + add
		} else {
			t = s1[len1-i].Val + add
		}
		add = t / 10
		node := &ListNode{
			t % 10,
			head.Next,
		}
		head.Next = node
	}
	if add > 0 {
		node := &ListNode{
			add,
			head.Next,
		}
		head.Next = node
	}
	return head.Next
}

func main() {
	t := addTwoNumbers(buildList([]int{7, 2, 4, 3}), buildList([]int{5, 6, 4}))
	for ; t != nil; t = t.Next {
		fmt.Println(t.Val)
	}
}
