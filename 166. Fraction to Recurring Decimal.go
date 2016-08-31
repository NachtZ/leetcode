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
func fractionToDecimal(numerator int, denominator int) string {
	ret := ""
    if numerator <0 && denominator>0||numerator >0 && denominator<0{
		ret = "-"
	} 
	n,d := int(math.Abs(float64(numerator))),int(math.Abs(float64(denominator)))
	res := n%d
	ret += strconv.Itoa(n/d)
	if res == 0{
		return ret
	}else{
		ret += "."
	}
	m := make(map[int]int)
	for res != 0{
		if _,ok := m[res];ok == true{
			ret = ret[0:m[res]] + "(" + ret[m[res]:]
			ret += ")"
			break
		}
		m[res] = len(ret)
		res *=10
		ret += strconv.Itoa(res/d)
		res %= d
	}
	return ret
}
func main() {
	fmt.Println(fractionToDecimal(2,3))
}
