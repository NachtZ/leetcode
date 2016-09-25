<<<<<<< HEAD
package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
	"unicode"
	//"math/rand"
	//"sort"
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

func max(a,b int)int{
	if a<b{
		return b
	}
	return a
}
func min(a,b int)int{
	if a>b{
		return b
	}
	return a
}
/**************************************************************/
var res string
func help(num string, k int) {
	fmt.Println("1:",num,"|",res,"|",k)
=======
var res string
func help(num string, k int) {
>>>>>>> a38f5365aa7083d4f4bb5157c44dcece331c9228
    if k <= 0{
		res = res + num
		return
	}
	if k >= len(num){
		return ;
	}
	minIdx := 0
	for i:=1;i<=k;i++{
		if num[i] < num[minIdx]{
			minIdx = i
		}
	}
<<<<<<< HEAD
	fmt.Println(num,"|",minIdx,"|",res,"|",k)
=======

>>>>>>> a38f5365aa7083d4f4bb5157c44dcece331c9228
	res = res + string(num[minIdx])
	new := num[minIdx+1:]
	help(new,k - minIdx)

}

func removeKdigits(num string, k int) string {
    res = ""
	help(num,k)
	i:=0
	for i=0;i<len(res)&&res[i]=='0';i++{
	}
	if i<len(res){
		res = res[i:]
	}else{
		res = "0"
	}
	return res
<<<<<<< HEAD
}

func main() {
	fmt.Println(removeKdigits("1432219",3))
}
=======
}
>>>>>>> a38f5365aa7083d4f4bb5157c44dcece331c9228
