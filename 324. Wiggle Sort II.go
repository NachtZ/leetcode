package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
	"unicode"
	//"container/heap"
	"sort"
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
type ss []int
func(n ss)Less(i,j int)bool{
	return n[i]<n[j]
}
func(n ss)Swap(i,j int){
	n[i],n[j] = n[j],n[i]
}
func (n ss)Len()int{
	return len(n)
}
func wiggleSort(nums []int){
	t := make(ss,len(nums))
	for i:=0;i<len(t);i++{
		t[i] = nums[i]
	}
	sort.Sort(t)
	j:=len(nums) -1
	for i:=1;i<len(nums);i,j=i+2,j-1{
		nums[i] = t[j]
	}
	for i:=0;i<len(nums);i,j=i+2,j-1{
		nums[i] = t[j]
	}
}

func main() {
	num := []int{1,1,2,1,2,2,1}
	wiggleSort(num)
	fmt.Println(num)
}
