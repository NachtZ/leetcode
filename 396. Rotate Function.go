<<<<<<< HEAD
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
=======
>>>>>>> a38f5365aa7083d4f4bb5157c44dcece331c9228
func maxRotateFunction(A []int) int {
    sum,f:= 0,0
	for i:=0;i<len(A);i++{
		sum += A[i]
		f += i*A[i]
	}
	m := f
	for i:=1;i<len(A);i++{
		t :=f+sum - A[len(A)-i]*len(A)
		if m < t{
			m = t
		}
		f += sum - A[len(A)-i]*len(A)
	}
	return m
<<<<<<< HEAD
}
func main() {
	fmt.Println(maxRotateFunction([]int{4, 3, 2, 6}))
	//fmt.Println([][]int{},0)
}
=======
}
>>>>>>> a38f5365aa7083d4f4bb5157c44dcece331c9228
