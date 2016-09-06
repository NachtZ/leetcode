package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
	"unicode"
	//"container/heap"
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

func min(a,b int)int{
	if a>b{
		return b
	}
	return a
}

func max(a,b int)int{
	if a<b{
		return b
	}
	return a
}

/**************************************************************/
func cau(a,b int, op byte)int{
	if op == '+'{
		return a+b
	}
	if op == '-'{
		return a-b
	}
	return a*b
}
func diffWaysToCompute(input string) []int {
    ret:= []int{}
	for i:=0;i<len(input);i++{
		if input[i] == '+'||input[i] == '-'||input[i] == '*'{
			left := diffWaysToCompute(input[:i])
			right := diffWaysToCompute(input[i+1:])
			for j:=0;j<len(left);j++{
				for k:=0;k<len(right);k++{
					ret = append(ret,cau(left[j],right[k],input[i]))
				}
			}
		}
	}
	if len(ret) == 0{
		v,_ := strconv.Atoi(input)
		ret = append(ret,v)
	}
	return ret
}
func main() {
	fmt.Println(diffWaysToCompute("2-1-1"))
}
