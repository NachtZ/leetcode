package main

import "fmt"

//import "unicode"
//import "strings"
import "strconv"
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
func evalRPN(tokens []string) int {
    stack := []int{}
	for _,v := range tokens{
		switch v {
		case "+":
			t := stack[len(stack)-2] + stack[len(stack)-1]
			stack = append(stack[:len(stack)-2],t)
		case "-":
			t := stack[len(stack)-2] - stack[len(stack)-1]
			stack = append(stack[:len(stack)-2],t)			
		case "*":
			t := stack[len(stack)-2] * stack[len(stack)-1]
			stack = append(stack[:len(stack)-2],t)
		case "/":
			t := stack[len(stack)-2] / stack[len(stack)-1]
			stack = append(stack[:len(stack)-2],t)
		default:
			t,_ := strconv.Atoi(v)
			stack = append(stack,t)
		}
	}
	return stack[0]
}
func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Print("Hello")
}
