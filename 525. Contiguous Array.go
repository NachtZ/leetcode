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
func findMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([][]int, len(nums)+1)
	res := 0
	for i := 0; i <= len(nums); i++ {
		dp[i] = make([]int, 2)
	}
	maps := make(map[int]int)
	for i := 1; i <= len(nums); i++ {
		if nums[i-1] == 0 {
			dp[i][0] = dp[i-1][0] + 1
			dp[i][i-1] = dp[i-1][1]
		} else {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = dp[i-1][1] + 1
		}
		if dp[i][0] == dp[i][1] {
			res = max(res, dp[i][0]*2)
		} else {
			dif := dp[i][1] - dp[i][0]
			if _, ok := maps[dif]; ok {
				res = max(res, 2*(dp[i][0]-dp[maps[dif]][0]))
			} else {
				maps[dif] = i
			}
		}
	}
	return res
}

func main() {
	fmt.Println(findMaxLength([]int{0, 1, 0}))
}
