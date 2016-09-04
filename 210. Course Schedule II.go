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

func findOrder(numCourses int, prerequisites [][]int) []int {
	if numCourses == 0{
		return []int{}
	}
    m := make([][]int,numCourses)
	in := make([]int,numCourses)
	queue := []int{}
	order := []int{}

	for _,p := range prerequisites{//build a map
			m[p[0]] = append(m[p[0]],p[1])
			in[p[1]]++
	}
	for i:=0;i<numCourses;i++{
		if in[i] == 0{
			queue = append(queue,i)
		}
	}
	for len(queue) != 0{
		curCourse := queue[0]
		queue = queue[1:]
		order = append(order,curCourse)
		for i:=0;i<len(m[curCourse]);i++{
			depender := m[curCourse][i]
			in[depender]--
			if in[depender] == 0{
				queue = append(queue,depender)
			}
		}
	}
	if len(order)!=numCourses{
		return []int{}
	}
	for i,j:=0,len(order)-1;i<j;i,j =i+1,j-1{
		order[i],order[j] = order[j],order[i]
	}
	return order

}

func main() {
	fmt.Println(findOrder(7,[][]int{}))
}
