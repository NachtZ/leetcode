package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
	"unicode"
	//"container/heap"
//	"sort"
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

func min(a,b int)int{
	if a>b{
		return b
	}
	return a
}

func max(a,b int)int{
	if a<b{
		return b
	}
	return a
}

/**************************************************************/
func min3(a,b,c int)(int){
	if a<b&&a<c{
		return a
	}
	if b<c{
		return b
	}
	return c
}
func nthUglyNumber(n int) int {
    t:=[]int{1}
	idx := make([]int,6)
	for i:=0;i<n;i++{
		a,b,c := t[idx[2]]*2,t[idx[3]]*3,t[idx[5]]*5
		v := min3(a,b,c)
		t = append(t,v)
		if v == a{
			idx[2]++
		}
		if v == b{
			idx[3]++
		}
		if v == c{
			idx[5]++
		}
	}
	return t[n-1]
}
func main() {
	fmt.Println(nthUglyNumber(10))
}
