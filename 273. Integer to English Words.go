package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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

func useLib() {
	fmt.Println(strconv.Itoa(1))
	fmt.Println(strings.Compare("1", "2"))
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
func help(a int) string {
	one := [20]string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
		"Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	ten := [10]string{"Zero", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	ret := ""
	h, t, o := a/100, (a%100)/10, a%10
	if h > 0 {
		ret += one[h] + " Hundred"
		if t > 0 || o > 0 {
			ret += " "
		}
	}
	if t > 0 {
		if t == 1 {
			ret += one[a%100]
		} else if o != 0 {
			ret += ten[t] + " " + one[o]
		} else {
			ret += ten[t]
		}
	} else if o > 0 {
		ret += one[o]
	}
	return ret
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	thr := [4]string{" Billion", " Million", " Thousand", ""}
	ret := ""
	count, t := 3, 0
	for num > 0 {
		t, num = num%1000, num/1000
		if t == 0 {
			count--
			continue
		}
		if len(ret) > 0 {
			ret = help(t) + thr[count] + " " + ret
		} else {
			ret = help(t) + thr[count]
		}
		count--
	}
	return ret
}

func main() {
	fmt.Println(numberToWords(1000000))
}
