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

func help(i,j,k int)(int, int){
	if i >j {
		i,j = j,i
	}
	if j > k{
		j,k = k,j
	}
	if i > j{
		i,j = j,i
	}
	return i,k
}

/**************************************************************/
func maxProduct(nums []int) int {
    max,min := make([]int,len(nums)),make([]int,len(nums))
	max[0],min[0] = nums[0],nums[0]
	max1 := nums[0]
	for i:=1;i<len(nums);i++{
		t1,t2 := max[i-1]*nums[i],min[i-1]*nums[i]
		min[i],max[i] = help(t1,t2,nums[i])
		if max[i]>max1{
			max1 = max[i]
		}
	}
	return max1

}
func main() {
	fmt.Println(maxProduct([]int{2,3,-2,4}))
	fmt.Print("Hello")
}
