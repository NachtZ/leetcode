package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
	"unicode"
	//"math/rand"
	//"sort"
	//"regexp"
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

func originalDigits(s string) string {
    nums := []string{"zero","one","two","three","four","five","six","seven","eight","nine"}
	var count [10]int
	var acount [26]int
	for i:=0;i<len(s);i++{
		acount[int(s[i]-'a')] ++
	}
	if acount[25]>0{
		str := nums[0]
		count[0] = acount[25]
		for i:=0;i<len(str);i++{
			acount[int(str[i]-'a')] -= count[0]
		}
	}
	if (acount[int('x'-'a')]>0){
		str := nums[6]
		count[6] = acount[int('x'-'a')]
		for i:=0;i<len(str);i++{
			acount[int(str[i]-'a')] -= count[6]
		}
	}
	if (acount[int('w'-'a')]>0){
		str := nums[2]
		count[2] = acount[int('w'-'a')]
		for i:=0;i<len(str);i++{
			acount[int(str[i]-'a')] -= count[2]
		}
	}

	if (acount[int('s'-'a')]>0){
		str := nums[7]
		count[7] = acount[int('s'-'a')]
		for i:=0;i<len(str);i++{
			acount[int(str[i]-'a')] -= count[7]
		}
	}
	if (acount[int('u'-'a')]>0){
		str := nums[4]
		count[4] = acount[int('u'-'a')]
		for i:=0;i<len(str);i++{
			acount[int(str[i]-'a')] -= count[4]
		}
	}
	if (acount[int('v'-'a')]>0){
		str := nums[5]
		count[5] = acount[int('v'-'a')]
		for i:=0;i<len(str);i++{
			acount[int(str[i]-'a')] -= count[5]
		}
	}
		if (acount[int('o'-'a')]>0){
		str := nums[1]
		count[1] = acount[int('o'-'a')]
		for i:=0;i<len(str);i++{
			acount[int(str[i]-'a')] -= count[1]
		}
	}	
	if (acount[int('n'-'a')]>0){
		str := nums[9]
		count[9] = acount[int('n'-'a')]/2
		for i:=0;i<len(str);i++{
			acount[int(str[i]-'a')] -= count[9]
		}
	}
	if (acount[int('i'-'a')]>0){
		str := nums[8]
		count[8] = acount[int('i'-'a')]
		for i:=0;i<len(str);i++{
			acount[int(str[i]-'a')] -= count[8]
		}
	}
	count[3] = acount[int('r'-'a')]
	str := ""
	for i:=0;i<=9;i++{
		for j:=0;j<count[i];j++{
			str += string(i+'0')
		}
	}
	return str
}
//[[12,13,1,12],[13,4,13,12],[13,8,10,12],[12,13,12,12],[13,13,13,13]]
func main() {

	fmt.Println(originalDigits("nine"))
}
