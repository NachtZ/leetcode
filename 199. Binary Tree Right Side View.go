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
func rightSideView(root *TreeNode) []int {
    ch := []*TreeNode{root}
	ret := []int{}
	if root == nil{
		return ret
	}
	right := root
	flag := root
	tmp := root
	for len(ch)!=0{
		tmp = ch[0]
		ch = ch[1:]
		if tmp.Left !=nil{
			right = tmp.Left
			ch = append(ch,tmp.Left)
		}
		if tmp.Right !=nil{
			right = tmp.Right
			ch = append(ch,tmp.Right)
		}
		if tmp == flag{
			ret = append(ret,tmp.Val)
			flag = right
		}
	}
	return ret
}
func main() {
	root := buildTree([]int{1,2,3,4})
	fmt.Println(rightSideView(root))
}
