package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
	"unicode"
	//"math/rand"
	"sort"
	//"regexp"
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

 type Interval struct {
 	   Start int
 	   End   int
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

type IntSlice []int

func (s IntSlice)Less(i,j int)bool{
	return s[i]<s[j]
}
func (s IntSlice)Swap(i,j int){
	s[i],s[j] = s[j],s[i]
}
func (s IntSlice)Len()int{
	return len(s)
}

/**************************************************************/
var count int
func dfs(A [][]int, idx int, sum int) int{
	if idx == 3{
		sum = -sum
		tmp := 0
		for i:=0;i<len(A[3]);i++{
			if sum < A[3][i]{
				return tmp 
			}
			if sum == A[3][i]{
				tmp ++ 
			}
		}
		return tmp
	}
	total := 0
	for i:=0;i<len(A[idx]);i++{
		tmp := dfs(A,idx+1,sum+A[idx][i])
		if tmp >0{
			total += tmp
			for ;i+1 < len(A[idx]) && A[idx][i+1] == A[idx][i];i++ {
				total += tmp
			}
		}
	}
	if idx == 0{
		count += total
	}
	return total
	

}
func fourSumCount(A []int, B []int, C []int, D []int) int {
    sort.Sort(IntSlice(A))
	sort.Sort(IntSlice(B))
	sort.Sort(IntSlice(C))
	sort.Sort(IntSlice(D))
	total := [][]int{A,B,C,D}
	count = 0
	//for i:=0;i<len(A);i++{
		dfs(total,0,0)
	//}
	return count
}
func main() {
	//fmt.Println(findContentChildren([]int{1,2},[]int{1,2,3}))	
	//fmt.Println(pathSum(buildTree([]int{1,2}),2))
	//fmt.Println(findKthNumber(13,2))
	fmt.Println(fourSumCount([]int{1,2},[]int{-2,-1},[]int{-1,2},[]int{0,2}))
}