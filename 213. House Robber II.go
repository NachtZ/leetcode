package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
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

func rob(nums []int) int {
	n := len(nums)
	if n == 0{
		return 0
	}
	if n ==1{
		return nums[0]
	}
	if n == 2{
		if nums[0]> nums[1]{
			return nums[0]
		}
		return nums[1]
	}
	if n == 3{
		if nums[0]>= nums[1] && nums[0]>=nums[2]{
			return nums[0]
		}
		if nums[1]>=nums[2]{
			return nums[1]
		}
		return nums[2]
	}
	max := nums[0]
	dp := make([]int,n)
	dp[0],dp[1],dp[2] = 0,nums[1],nums[2]
	for i:=3;i<n;i++{
		if dp[i-2]>dp[i-3]{
			dp[i] = dp[i-2]+nums[i]
		}else {
			dp[i] = dp[i-3] + nums[i]
		}
		if dp[i]>max{
			max = dp[i]
		}
	}
	dp[0],dp[1],dp[2] = nums[0],0,nums[0]+nums[2]
	for i:=3;i<n-1;i++{
		if dp[i-2]>dp[i-3]{
			dp[i] = dp[i-2]+nums[i]
		}else {
			dp[i] = dp[i-3] + nums[i]
		}	
		if dp[i]>max{
			max = dp[i]
		}	
	}
	return max
}
func main() {
	fmt.Println(rob([]int{1,1,1}))
}
