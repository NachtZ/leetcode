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
func dfs(grid [][]byte,x,y int){
	if x <0 || x>=len(grid) || y<0 || y>=len(grid[0]){
		return
	}
	if grid[x][y] == '1'{
		grid[x][y] = '2'
		dfs(grid,x+1,y)
		dfs(grid,x-1,y)
		dfs(grid,x,y+1)
		dfs(grid,x,y-1)
	}
}
func numIslands(grid [][]byte) int {
	if len(grid) == 0{
		return 0
	}
    m,n := len(grid),len(grid[0])
	count :=0
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if grid[i][j] == '1'{
				count ++
				dfs(grid,i,j)
			}
		}
	}
	return count
}
func main() {
	fmt.Println(numIslands([][]byte{[]byte("1"),[]byte("1")}))
}
