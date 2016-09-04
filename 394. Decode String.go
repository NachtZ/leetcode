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

func decodeString(s string) string {
    times := 0
	ret := ""
	for i:=0;i<len(s);i++{
		if s[i] == '['{
			t := 1
			j:=i+1
			for j=i+1;t!=0&& j<len(s);j++{
				if s[j] == '['{
					t++
				}else if s[j] == ']'{
					t --
				}
			}
			tmp := s[i+1:j]
			i = j-1
			str := decodeString(tmp)
			for times >0{
				ret += str
				times --
			}
		}else if s[i]<='9'&&s[i]>='0'{
			times = 10*times + int(s[i]-'0')
		}else if s[i]!=']'{
			ret += string(s[i])
		}
	}
	return ret
}

func main() {
	fmt.Println(decodeString("2[abc]3[cd]ef"))
}
