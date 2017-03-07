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
func updateBoard(board [][]byte, click []int) [][]byte {
	x, y := click[0], click[1]
	move := [][]int{
		[]int{1, 0},
		[]int{-1, 0},
		[]int{0, 1},
		[]int{0, -1},
		[]int{1, 1},
		[]int{1, -1},
		[]int{-1, 1},
		[]int{-1, -1},
	}
	switch board[x][y] {
	case 'M':
		board[x][y] = 'X'
	case 'E':
		count := 0
		for i := 0; i < 8; i++ {
			newx, newy := x+move[i][0], y+move[i][1]
			if newx < 0 || newx >= len(board) || newy < 0 || newy >= len(board[0]) {
				continue
			}
			if board[newx][newy] == 'M' {
				count++
			}
		}
		if count > 0 {
			board[x][y] = byte('0' + count)
		} else {
			board[x][y] = 'B'
			for i := 0; i < 8; i++ {
				newx, newy := x+move[i][0], y+move[i][1]
				if newx < 0 || newx >= len(board) || newy < 0 || newy >= len(board[0]) || board[newx][newy] != 'E' {
					continue
				}
				updateBoard(board, []int{newx, newy})
			}
		}
	}
	return board
}
func main() {
	b := [][]byte{
		[]byte{'E', 'E', 'E', 'E', 'E'},
		[]byte{'E', 'E', 'M', 'E', 'E'},
		[]byte{'E', 'E', 'E', 'E', 'E'},
		[]byte{'E', 'E', 'E', 'E', 'E'},
	}
	fmt.Println(updateBoard(b, []int{3, 0}))
}
