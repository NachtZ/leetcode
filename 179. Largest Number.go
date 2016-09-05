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
func help(nums []string, idx int) string {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 0 {
		return ""
	}
	if len(nums) == 2 {
		if nums[0]+nums[1] > nums[1]+nums[0] {
			return nums[0] + nums[1]
		} else {
			return nums[1] + nums[0]
		}
	}
	n := 0
	bucket := make([][]string, 11)
	for i := 0; i < 11; i++ {
		bucket[i] = []string{}
	}
	for _, v := range nums {
		if len(v) < idx {
			bucket[10] = append(bucket[10], v)
		} else {
			if len(v) == idx {
				n = 0
			} else {
				n = int(v[idx] - '0')
			}
			bucket[9-n] = append(bucket[9-n], v)
		}
	}
	ret := ""
	for i := 0; i < 10; i++ {
		if len(bucket[i]) != 0 {
			ret += help(bucket[i], idx+1)
		}
	}
	tmp := ""
	for i := 0; i < len(bucket[10]); i++ {
		tmp += bucket[10][i]
	}
	if len(tmp) > 0 {
		if tmp+ret > ret+tmp {
			ret = tmp + ret
		} else {
			ret += tmp
		}
	}
	return ret
}

func largestNumber(nums []int) string {
	str := make([]string, len(nums))
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
		str[i] = strconv.Itoa(nums[i])
	}
	if total == 0 {
		return "0"
	}
	return help(str, 0)
}
func main() {
	fmt.Println(largestNumber([]int{824, 938, 1399, 5607, 6973, 5703, 9609, 4398, 8247}))
}
