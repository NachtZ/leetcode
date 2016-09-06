package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
	"unicode"
//	"container/heap"
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
func majorityElement(nums []int) []int {
    if len(nums)<=1{
		return nums
	}
	num1,num2 := nums[0],0
	c1,c2 :=1,0
	for i:=1;i<len(nums);i++{
		if nums[i] == num1{
			c1 ++
		}else if nums[i] == num2{
			c2 ++
		}else if c1 <=0{
			num1 = nums[i]
			c1 = 1
		}else if c2 <= 0{
			num2 = nums[i]
			c2 = 1
		}else{
			c1--
			c2--
		}
	}
	c1,c2 = 0,0
	for i:=0;i<len(nums);i++{
		if nums[i] == num1{
			c1++
		}else if nums[i] == num2{
			c2 ++
		}
	}
	ret := []int{}
	if c1 > len(nums)/3{
		ret = append(ret,num1)
	}
	if c2 > len(nums)/3{
		ret = append(ret,num2)
	}
	return ret
}

func main() {
	fmt.Println(majorityElement([]int{1,1,1,1,1,2,2,2,2,2,2}))
}
