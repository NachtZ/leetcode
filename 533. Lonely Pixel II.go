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

func findBlackPixel(picture [][]byte, N int) int {
	if len(picture) == 0 || len(picture[0]) == 0 {
		return 0
	}
	dic := make(map[string]int)
	strs := make([]string, len(picture))
	arrC := make([]int, len(picture[0]))
	arrR := make([]int, len(picture))
	for i := 0; i < len(picture); i++ {
		str := ""
		for j := 0; j < len(picture[0]); j++ {
			str += string(picture[i][j])
			if picture[i][j] == 'B' {
				arrC[j]++
				arrR[i]++
			}
		}
		strs[i] = str
		dic[str]++
	}
	count := 0
	for i := 0; i < len(picture); i++ {
		if dic[strs[i]] != N {
			continue
		}
		for j := 0; j < len(picture[0]); j++ {
			if picture[i][j] == 'B' {
				if arrC[j] == arrR[i] && arrC[j] == N {
					count++
				}
			}
		}
	}
	return count
}
func main() {
	pic := make([][]byte, 6)
	str := []string{"WBWBBW", "BWBWWB", "WBWBBW", "BWBWWB", "WWWBBW", "BWBWWB"}
	for i := 0; i < len(str); i++ {
		pic[i] = make([]byte, 6)
		for j := 0; j < len(str[i]); j++ {
			pic[i][j] = str[i][j]
		}
	}
	fmt.Println(findBlackPixel(pic, 3))

}
