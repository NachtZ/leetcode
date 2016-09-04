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
var ret [][]int
var path []int
func help(k,idx, n int){
	if idx == n && k == 1{
		tmp := []int{}
		tmp = append(tmp,path...)
		tmp = append(tmp,idx)
		ret = append(ret,tmp)
		return
	}
	if k == 1{
		return
	}
	path = append(path,idx)
	n = n-idx
	for i := idx+1;i<=9;i++{
		help(k-1,i,n)
	}
	path = path[:len(path)-1]
}
func combinationSum3(k int, n int) [][]int {
    ret = [][]int{}
	for i:=1;i<=10-k;i++{
		path = []int{} 
		help(k,i,n)
	}
	return ret
}
func main() {
	fmt.Println(combinationSum3(3,9))
}
