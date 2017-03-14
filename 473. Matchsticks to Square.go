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
*/
/**************************************************************/
/*time out
func countBit1(num int) int {
	count := 0
	for num > 0 {
		count++
		num = num & (num - 1)
	}
	return count
}
func totalHammingDistance(nums []int) int {
	ret := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			ret += countBit1(nums[i] ^ nums[j])
		}
	}
	return ret
}
*/

func dfs(nums []int, used int, length int, left int, m map[[2]int]bool) bool {
	if _, ok := m[[2]int{used, left}]; ok {
		return false
	}
	if left == 0 && used == 0 {
		return true
	}
	t := 1
	if left == 0 {
		left = length
	}
	for i := 0; i < len(nums); i++ {
		if used&t == 0 {
			t <<= 1
			continue
		}
		if left >= nums[i] {
			t := dfs(nums, used&(^t), length, left-nums[i], m)
			if t {
				return true
			}
		}
		t <<= 1
	}
	m[[2]int{used, left}] = false
	return false
}

func makesquare(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	side := 0
	t := 0
	for i := 0; i < len(nums); i++ {
		t <<= 1
		t |= 1
		side += nums[i]
	}
	if side%4 > 0 {
		return false
	}
	m := make(map[[2]int]bool)
	return dfs(nums, t, side/4, side/4, m)

}
func main() {

	fmt.Println(makesquare([]int{5, 5, 5, 5, 16, 4, 4, 4, 4, 4, 3, 3, 3, 3, 4}))
}
