package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
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
	res :=0
	var op byte
	op = '+'
	for i:=0;i<len(s);i++{
		if s[i] == '('{
			t ,j:= 1,i+1
			for ;j<len(s)&&t>0;j++{
				if s[j] == '('{
					t++
				}else if s[j] == ')'{
					t--
				}
			}
			num = calculate(s[i+1:j-1])
			i = j-1
		}else if s[i] == ' '||s[i] == ')'{
			continue
		}else if s[i] == '+' || s[i] == '-'{
			if op == '+'{
				res += num
			}else if op == '-'{
				res -= num
			}
			op = s[i]
			num = 0
		}else{
			num = num*10 + int(s[i]-'0')
		}
	}
	if op == '+'{
			res += num
	}else if op == '-'{
			res -= num
	}	
	return res	
}

func main() {
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
}
