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
var count, TOTAL int
var sup = [][]int{
	[]int{},
	[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
	[]int{1, 2, 4, 6, 8, 10, 12, 14},
	[]int{1, 3, 6, 9, 12, 15},
	[]int{1, 2, 4, 8, 12},
	[]int{1, 5, 10, 15},
	[]int{1, 2, 3, 6, 12},
	[]int{1, 7, 14},
	[]int{1, 2, 4, 8},
	[]int{1, 3, 9},
	[]int{1, 2, 5, 10},
	[]int{1, 11},
	[]int{1, 2, 3, 4, 6, 12},
	[]int{1, 13},
	[]int{1, 2, 7, 14},
	[]int{1, 3, 5, 15},
}

var used [16]bool

func help(idx int) {
	if idx > TOTAL {
		count++
		return
	}
	for _, v := range sup[idx] {
		if v <= TOTAL && used[v] == false {
			used[v] = true
			help(idx + 1)
			used[v] = false
		}
	}

}

func countArrangement(N int) int {
	for i := 0; i < len(used); i++ {
		used[i] = false
	}
	count = 0
	TOTAL = N
	help(1)
	return count
}

func main() {
	fmt.Println(countArrangement(2))
}
