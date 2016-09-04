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

func divid(s string, sp map[byte]bool)[]string{
	ret := []string{}
	start := 0
	for i:=0;i<len(s);i++{
		if _,ok := sp[s[i]];ok{
			tmp := s[start:i]
			ret = append(ret,tmp)
			start = i+1
		}
	}
	if start <len(s){
		tmp := s[start:]
		ret = append(ret,tmp)
	}
	return ret
}

func longestSubstring(s string, k int) int {
	m := make(map[byte]int)
	found := false
	for i:=0;i<len(s);i++{
		_,ok := m[s[i]]
		if ok{
			m[s[i]]++
		}else{
			m[s[i]] = 1
		}
	}
	splites := make(map[byte]bool)
	for key,v := range m{
		if v < k{
			splites[key] = true
			found = true
		}
	}
	if found == false{
		return len(s)
	}
	ret := divid(s,splites)
	max := 0
	for i:=0;i<len(ret);i++{
		val := longestSubstring(ret[i],k)
		if val > max{
			max = val
		}
	}
	return max
}

func main() {
	fmt.Println(longestSubstring("aaabb",3))
}
