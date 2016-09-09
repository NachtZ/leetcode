package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
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

func useLib() {
	fmt.Println(strconv.Itoa(1))
	fmt.Println(strings.Compare("1", "2"))
	fmt.Println(math.Abs(1.0))
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

/**************************************************************/
type env [][]int

func (e env) Len() int {
	return len(e)
}
func (e env) Less(a, b int) bool {
	if e[a][0] == e[b][0] {
		return e[a][1] < e[b][1]
	}
	return e[a][0] < e[b][0]
}
func (e env) Swap(a, b int) {
	e[a][0], e[b][0] = e[b][0], e[a][0]
	e[a][1], e[b][1] = e[b][1], e[a][1]
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}
	sort.Sort(env(envelopes))
	dp := make([]int, len(envelopes))
	max := 1
	dp[0] = 1
	for i := 1; i < len(envelopes); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if envelopes[i][0] > envelopes[j][0] && envelopes[i][1] > envelopes[j][1] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}
func main() {
	fmt.Println(maxEnvelopes([][]int{}))
}
