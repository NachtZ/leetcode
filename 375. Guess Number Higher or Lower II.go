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
func getMoneyAmount(n int) int {
	dp := make([][]int,n+2)
	for i:=0;i<n+2;i++{
		dp[i] = make([]int,n+2)
	}
	for l:=1;l<n;l++{
		for i :=1;i<=n-l;i++{
			j := i+l
			dp[i][j] = 655350
			for k:=i;k<=j;k++{
				dp[i][j] = min(dp[i][j],k+max(dp[i][k-1],dp[k+1][j]))
			}
		}
	}
	return dp[1][n]
}
func main() {
	fmt.Println(getMoneyAmount(10))
	//fmt.Println([][]int{},0)
}
