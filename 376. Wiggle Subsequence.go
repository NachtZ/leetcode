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
func wiggleMaxLength(nums []int) int {
    if len(nums)<2{
        return len(nums)
    }
	dp := make([][2]int,len(nums))
	maxs := 1
	//idx := 0
	dp[0][1],dp[0][0] = 1, 1//0 for +, 1 for -
    for i:=1;i<len(nums);i++{
		dp[i][0],dp[i][1] = 1,1
	//	idx = 1- idx
		for j :=i-1;j>=0;j--{
			if nums[i] == nums[j]{
				continue
			}
			if nums[i] > nums[j]{
				dp[i][0] = max(dp[i][0],dp[j][1]+1)
			}else{
				dp[i][1] = max(dp[i][1],dp[j][0]+1)
			}
		}
		if maxs < dp[i][0]{
			maxs = dp[i][0]
		}
		if maxs < dp[i][1]{
			maxs = dp[i][1]
		}
	}
	return maxs
}
func main() {
	fmt.Println(wiggleMaxLength([]int{1,7,4,9,2,5}))
	//fmt.Println([][]int{},0)
}
