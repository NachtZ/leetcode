package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
	"unicode"
	//"math/rand"
	//"sort"
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

func useLib(){
	fmt.Println(strconv.Itoa(1))
	fmt.Println(strings.Compare("1","2"))
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



func max(a,b int)int{
	if a<b{
		return b
	}
	return a
}
func min(a,b int)int{
	if a>b{
		return b
	}
	return a
}
/**************************************************************/
func isRectangleCover(rectangles [][]int) bool {
	if len(rectangles) == 1{
		return true
	}
	m := make(map[[2]int]int)
	var lt,ld,rt,rd [2]int
	corner:= [4]int{65535,65535,-65535,-65535}
	area := 0
	for i:=0;i<len(rectangles);i++{
		lt = [2]int{rectangles[i][0],rectangles[i][3]}
		ld = [2]int{rectangles[i][0],rectangles[i][1]}
		rt = [2]int{rectangles[i][2],rectangles[i][3]}
		rd = [2]int{rectangles[i][2],rectangles[i][1]}
		corner[0] = min(corner[0],rectangles[i][0])
		corner[1] = min(corner[1],rectangles[i][1])
		corner[2] = max(corner[2],rectangles[i][2])
		corner[3] = max(corner[3],rectangles[i][3])
		area += (rectangles[i][2]-rectangles[i][0])*(rectangles[i][3]-rectangles[i][1])
		m[lt],m[ld],m[rt],m[rd] = m[lt]+1,m[ld]+1,m[rt]+1,m[rd]+1
	}
	if area != (corner[2]-corner[0])*(corner[3]-corner[1]){
		return false
	}
	fmt.Println(corner)
	ltc := [2]int{corner[0],corner[3]}
	ldc := [2]int{corner[0],corner[1]}
	rtc := [2]int{corner[2],corner[3]}
	rdc := [2]int{corner[2],corner[1]}
	if m[ldc]!=1 || m[rtc]!=1||m[rtc] !=1 || m[rdc]!=1{
		return false
	}
	for k,v := range m{
		if ldc==k || rtc==k||rdc ==k || ltc==k{
			continue
		}
		if v != 2 && v!= 4{
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(isRectangleCover([][]int{[]int{0,0,3,3},[]int{1,1,2,2},[]int{1,1,2,2}}))
	//fmt.Println([][]int{},0)
}
