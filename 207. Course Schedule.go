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
func help(idx int, res []int, m map[int][]int) bool{
	if res[idx] == 1{
		return true
	}
	if res[idx] == 2{
		return false
	}
	v,ok := m[idx]
	if ok == false{
		return true
	}
	res[idx] = 2
	for i:=0;i<len(v);i++{
		if help(v[i],res,m) == false{
			return false
		}
	}
	res[idx] =1 
	return true
}
func canFinish(numCourses int, prerequisites [][]int) bool {
    m := make(map[int][]int)
	res := make([]int,numCourses)//1 for can finish, 0 for unknow, 2 for is visited
	for _,p := range prerequisites{//build a map
		v,ok := m[p[0]]
		if ok == false{
			m[p[0]] = []int{p[1]}
		}else{
			m[p[0]] = append(v,p[1])
		}
	}
	for i:=0;i<numCourses;i++{
		if help(i,res,m)== false{
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(canFinish(2,[][]int{[]int{1,0},[]int{0,1}}))
}
