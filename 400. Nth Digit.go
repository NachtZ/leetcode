<<<<<<< HEAD
=======
package main

import (
	"fmt"
	"math"
	//"sort"
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
>>>>>>> a38f5365aa7083d4f4bb5157c44dcece331c9228
func findNthDigit(n int) int {
	i := 0
	mark, l := 1, 0
	total := 0
	for i = 1; total < n; i++ {
		if i >= mark {
			total += l + 1
			l++
			mark *= 10
		} else {
			total += l
		}
	}
	i--
	total -= l
	t := strconv.Itoa(i)
	return int(t[n-total-1] - '0')
<<<<<<< HEAD
}
=======
}
func main() {
	fmt.Println(findNthDigit(11))
}
>>>>>>> a38f5365aa7083d4f4bb5157c44dcece331c9228
