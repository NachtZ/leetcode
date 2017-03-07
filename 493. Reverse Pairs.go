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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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

type StrSlice []string

func (s StrSlice) Less(i, j int) bool {
	if len(s[i]) != len(s[j]) {
		return len(s[i]) > len(s[j])
	}
	return s[i] < s[j]
}
func (s StrSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s StrSlice) Len() int {
	return len(s)
}

*/
/**************************************************************/
func mergeAndsort(nums []int, left, mid, right int) int {
	ret, i, j := 0, 0, 0
	idx := 0
	tmp := make([]int, right-left+1)
	for i, j = left, mid+1; i <= mid && j <= right; {
		if nums[i] > 2*nums[j] {
			ret += mid - i + 1
			j++
		} else {
			i++
		}
	}
	for i, j = left, mid+1; i <= mid && j <= right; {
		if nums[i] < nums[j] {
			tmp[idx] = nums[i]
			i++
		} else {
			tmp[idx] = nums[j]
			j++
		}
		idx++
	}
	for i <= mid {
		tmp[idx] = nums[i]
		idx++
		i++
	}
	for j <= right {
		tmp[idx] = nums[j]
		idx++
		j++
	}
	for i, idx := left, 0; i <= right; i, idx = i+1, idx+1 {
		nums[i] = tmp[idx]
	}
	return ret
}
func help(nums []int, left, right int) int {
	if left >= right {
		return 0
	}
	ret := 0
	mid := left + (right-left)/2
	ret += help(nums, left, mid) + help(nums, mid+1, right)
	ret += mergeAndsort(nums, left, mid, right)
	return ret
}
func reversePairs(nums []int) int {
	return help(nums, 0, len(nums)-1)
}
func main() {

	fmt.Println(reversePairs([]int{2, 4, 3, 5, 1}))
}
