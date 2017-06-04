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
func printTreeNode(root *TreeNode) {
	if root == nil {
		return
	}
	left, right := "nil", "nil"
	if root.Left != nil {
		left = strconv.Itoa(root.Left.Val)
	}
	if root.Right != nil {
		right = strconv.Itoa(root.Right.Val)
	}
	fmt.Println(root.Val, ":", left, right)
	printTreeNode(root.Left)
	printTreeNode(root.Right)
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

type IntSlice []int

func (s IntSlice) Less(i, j int) bool {
	return s[i] > s[j]
}
func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s IntSlice) Len() int {
	return len(s)
}

/**************************************************************/
//helper to caculate that in now situation, can now player win.
//dp means at this situation, whether the player now can win.
func helper(maxInt int, left int, bit int, dp map[int]bool, bitmap []int) bool {
	if val, ok := dp[bit]; ok {
		return val
	}
	if left <= maxInt {
		for i := maxInt; i >= left; i-- {
			if (bit & bitmap[i]) == 0 {
				dp[bit] = true
				return true
			}
		}
	}
	for i := maxInt; i > 0; i-- {
		if (bit & bitmap[i]) > 0 {
			continue
		}
		bit |= bitmap[i]
		ret := helper(maxInt, left-i, bit, dp, bitmap)
		bit &= (^bitmap[i])
		if ret == false {
			dp[bit] = true
			return true
		}
	}
	dp[bit] = false
	return false
}
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	dp := make(map[int]bool)
	sum := 0
	bitmap := make([]int, maxChoosableInteger+1)
	for i := maxChoosableInteger; i > 0; i-- {
		sum += i
		bitmap[i] = 1 << uint8(i)
	}
	if sum < desiredTotal {
		return false
	}
	return helper(maxChoosableInteger, desiredTotal, 0, dp, bitmap)

}
func main() {
	fmt.Println(canIWin(5, 50))
}
