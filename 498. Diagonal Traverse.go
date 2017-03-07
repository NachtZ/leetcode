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
func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	flag := false
	ret := make([]int, 0, len(matrix)*len(matrix[0]))
	for t, v := 0, len(matrix)+len(matrix[0]); t < v; t++ {
		for i := 0; i <= t; i++ {
			if flag {
				if i < len(matrix) && t-i < len(matrix[0]) {
					ret = append(ret, matrix[i][t-i])
				}
			} else {
				if i < len(matrix[0]) && t-i < len(matrix) {
					ret = append(ret, matrix[t-i][i])
				}
			}
		}
		flag = !flag
	}
	return ret
}
func main() {
	//fmt.Println(findRotateSteps("godding", "godding"))

	fmt.Println(findDiagonalOrder([][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
		[]int{10, 11, 12},
	}))
}
