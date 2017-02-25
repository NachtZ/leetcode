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
var maps map[int]int
var maxVal []int

func calSubSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := root.Val
	if root.Left != nil {
		ret += calSubSum(root.Left)
	}
	if root.Right != nil {
		ret += calSubSum(root.Right)
	}
	root.Val = ret
	if _, ok := maps[ret]; ok {
		maps[ret]++
	} else {
		maps[ret] = 1
	}
	if maps[ret] == maxVal[0] {
		maxVal = append(maxVal, ret)
	} else {
		if maps[ret] > maxVal[0] {
			maxVal = []int{maps[ret], ret}
		}
	}

	return ret
}

func findFrequentTreeSum(root *TreeNode) []int {
	maps = make(map[int]int)
	maxVal = []int{0}
	maps[0] = 0
	calSubSum(root)
	set := make(map[int]bool)
	ret := make([]int, 0, len(maxVal)-1)
	for _, val := range maxVal[1:] {
		_, ok := set[val]
		if ok == false {
			set[val] = true
			ret = append(ret, val)
		}
	}
	return ret
}
func main() {
	fmt.Println(findFrequentTreeSum(buildTree([]int{5, 2, -5})))

}
