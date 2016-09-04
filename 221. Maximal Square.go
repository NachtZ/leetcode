package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
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

func min (a,b,c int)int{
	if a<=b && a <=c{
		return a
	}
	if b<=c{
		return b
	}
	return c
}

func maximalSquare(matrix [][]byte) int {
    if len(matrix) == 0 || len(matrix[0]) == 0{
		return 0
	}
	m,n := len(matrix),len(matrix[0])
	dp := make([][]int,m)
	for i:=0;i<m;i++{
		dp[i] = make([]int,n)
	}
	max :=0
	for i:=0;i<m;i++{
		dp[i][0] = int(matrix[i][0] - '0')
		if max < dp[i][0]{
			max = dp[i][0]
		}
	}
	for i:=0;i<n;i++{
		dp[0][i] = int(matrix[0][i] - '0')
		if max < dp[0][i]{
			max = dp[0][i]
		}
	}
	for i :=1;i<m;i++{
		for j:=1;j<n;j++{
			if matrix[i][j] == '0'{
				continue	
			}		
			dp[i][j] = min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1])+1
			if dp[i][j]>max{
				max = dp[i][j]
			}
		}
	}
	return max*max
}
func main() {
	fmt.Println(maximalSquare([][]byte{[]byte("10100"),[]byte("10111"),[]byte("11111"),[]byte("10010")}))
}
