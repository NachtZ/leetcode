package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
	"unicode"
//	"container/heap"
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
func calculate(s string) int {
    num :=0
	st := []int{}
	var sign byte
	sign = '+'
	for i:=0;i<len(s);i++{
		if unicode.IsDigit(rune(s[i])){
			num = num*10 + int(s[i]-'0')
		}
		if s[i]!=' '&&!unicode.IsDigit(rune(s[i]))||i==len(s)-1{
			switch sign{
				case '+':
				case '-':
					num = -num;
				case '*':
					num = st[len(st)-1]*num
					st = st[:len(st)-1]
				case '/':
					num = st[len(st)-1]/num
					st = st[:len(st)-1]
			}
			sign = s[i]
			st = append(st,num)
			num = 0
		}
	}
	for _,v := range st{
		num += v
	}
	return num
}

func main() {
	fmt.Println(calculate(" 3+5 / 2 "))
}
