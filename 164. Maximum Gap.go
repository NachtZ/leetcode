package main

import "fmt"

//import "unicode"
//import "strings"
//import "strconv"
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
func getM(nums []int) (max,min int){

	max,min = nums[0],nums[0]
	for i:=1;i<len(nums);i++{
		if max < nums[i]{
			max = nums[i]
		}
		if min>nums[i]{
			min = nums[i]
		}
	}
	return 
}
func maximumGap(nums []int) int {
    if len(nums)<2{
		return 0
	}
	max,min := getM(nums)
	if max == min{
		return 0
	}
	b := make([][]int, len(nums)+1)
	maxgap :=0
	for i:=0;i<len(nums);i++{
		idx := (nums[i]-min)*len(nums)/(max-min)
		b[idx] = append(b[idx],nums[i])
	}

	min,max = b[0][0],b[0][0]
	min1,max1 := b[0][0],b[0][0]
	for i:=0;i<len(b);i++{
		if len(b[i])>0{
			max1,min1 = getM(b[i])
			if min1 - max > maxgap{
				maxgap = min1 - max
			}
			if max1 - min1 > maxgap{
				maxgap = max1 - min1
			}
			max,min = max1,min1
		}
	}
	return maxgap

}
func main() {
	fmt.Println(maximumGap([]int{1,1,1,1}))
	fmt.Print("Hello")
}
