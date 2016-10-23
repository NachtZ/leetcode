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

func check(a,b,c [26]int) bool{
	for i:=0;i<26;i++{
		if a[i]-b[i] != c[i]{
			return false
		}
	}
	return true
}

func findAnagrams(s string, p string) []int {
	if len(s)<len(p){
		return []int{}
	}
	ret := []int{}
    start, end, mark := [26]int{},[26]int{},[26]int{}
	for i:=0;i<len(p);i++{
		mark[p[i]-'a'] ++
	}
	for i:=0;i<len(p);i++{
		end[s[i]-'a'] ++
	}
	if check(end,start,mark){
		ret = append(ret,0)
	}
	for i:=len(p);i<len(s);i++{
		start[s[i-len(p)]-'a'] ++
		end[s[i]-'a']++
		if check(end,start,mark){
			ret = append(ret,i-len(p)+1)
		}
	}
	return ret
}

//[[12,13,1,12],[13,4,13,12],[13,8,10,12],[12,13,12,12],[13,13,13,13]]
//"IMNJJTRMJEGMSOLSCCQICIHLQIOGBJAEHQOCRAJQMBIBATGLJDTBNCPIFRDLRIJHRABBJGQAOLIKRLHDRIGERENNMJSDSSMESSTR"
//2 6
//"AABABBA" 1 4
//"KJRGKSKKOKLPADMAGODEDNEBMJMKGAPNLSAKADRLHHDRMJTMFBSIQGHENKABISHQNRFJKEPMFIPNDNEOBRJEHFLIHMDLMCIHLHQN" 5 10
func main() {
	//fmt.Println(pathSum(buildTree([]int{10,5,-3,3,2,-1,11,3,-2,-1,1}),8))	
	//fmt.Println(pathSum(buildTree([]int{1,2}),2))
	fmt.Println(findAnagrams("abab","ab"))
}
