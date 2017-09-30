package main

import (
	"fmt"
	"math"
	//"sort"
	"strconv"
	"strings"
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

func useLib() {
	fmt.Println(strconv.Itoa(1))
	fmt.Println(strings.Compare("1", "2"))
	fmt.Println(math.Abs(1.0))
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

/**************************************************************/
func thirdMax(nums []int) int {
	max1, max2, max3 := -2147483648, -2147483648, -2147483648
	temp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		temp[nums[i]]++
		if max3 >= nums[i] || max1 == nums[i] || max2 == nums[i] {
			continue
		}
		max3 = nums[i]
		if max2 < max3 {
			max2, max3 = max3, max2
		}
		if max1 < max2 {
			max1, max2 = max2, max1
		}
	}
	if len(temp) <= 2 {
		return max1
	}
	return max3
}
func main() {
	fmt.Println(thirdMax([]int{1, 2, 3, 4, 5, 6}))
}
