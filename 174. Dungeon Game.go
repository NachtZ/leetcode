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
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}
func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}
func calculateMinimumHP(dungeon [][]int) int {
	m,n := len(dungeon),len(dungeon[0])
    dp := make([][]int,m)
	for i:=0;i<m;i++{
		dp[i] = make([]int,n)
	}
	dp[m-1][n-1] = max(0-dungeon[m-1][n-1],0)
	for i:=m-2;i>=0;i--{
		dp[i][n-1] =max(dp[i+1][n-1]-dungeon[i][n-1],0)
	}
	for i:=n-2;i>=0;i--{
		dp[m-1][i] =max(dp[m-1][i+1]-dungeon[m-1][i],0)
	}
	for i:=m-2;i>=0;i--{
		for j:=n-2;j>=0;j--{
			dp[i][j] = max(min(dp[i][j+1],dp[i+1][j])-dungeon[i][j],0)
		}
	}
	return dp[0][0]+1

}
func main() {
	fmt.Println(calculateMinimumHP([][]int{[]int{3,-20,30},[]int{-3,4,0}}))
}
