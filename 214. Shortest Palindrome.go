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
func shortestPalindrome(s string) string {
    tstr := "#"
	for _,v := range s{
		tstr += string(v)+"#"
	}
	p := make([]int,len(tstr)+1)
	mx,id :=0,0
	for i:=1;i<len(tstr);i++{
		if mx >i{
			if p[2*id-i]<mx-i{
				p[i] = p[2*id-i]
			}else{
				p[i] = mx -i
			}
		}else{
			p[i] = 1
		}
		for i-p[i]>0 && i+p[i]<len(tstr) && tstr[i-p[i]] == tstr[i+p[i]]{
			p[i]++
		}
		if i+p[i]>mx{
			id = i
			mx = i+p[i]
		}
	}
	mid, pos := len(tstr)/2, 0
	for i:=0;i<mid;i++{
		if p[mid-i] == mid -i{
			pos = mid-i
			break
		}
		if p[mid+i] == mid+i{
			pos = mid +i
			break
		}
	}
	if mid == pos{
		return s
	}
	t := ""
	for i := len(tstr)-1;i>=2*pos;i--{
		if tstr[i] !='#'{
			t += string(tstr[i])
		}
	}
	return t+s
}
func main() {
	fmt.Println(shortestPalindrome("ssts"))
}
