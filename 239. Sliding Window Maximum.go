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



func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0{
		return nums
	}
	left,right := make([]int,len(nums)),make([]int,len(nums))
	left[0],right[len(nums)-1] = nums[0],nums[len(nums)-1]
	for i,j:=1,len(nums)-2;i<len(nums);i,j = i+1,j-1{
		if i%k == 0{
			left[i] = nums[i]
		}else{
			left[i] = max(left[i-1],nums[i])
		}
		if j%k == 0{
			right[j] = nums[j]
		}else{
			right[j] = max(right[j+1],nums[j])
		}
	}
	ret := make([]int,len(nums)-k+1)
	for i:=0;i+k<=len(nums);i++{
		ret[i]=max(right[i],left[i+k-1])
	}
	return ret
}
func main() {
	fmt.Println(maxSlidingWindow([]int{2,1,3,4,6,3,8,9,10,12,56},4))
}
