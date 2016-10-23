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
func characterReplacement(s string, k int) int {
    dp := make([]int,len(s)+1)
	var ch [26]bool
	max := 0
	if k == 0{
		length := 1
		for i :=1;i<len(s);i++{
			if s[i-1] == s[i]{
				length ++
			}else{
				if max < length{
					max = length
				}
				length = 1
			}
			
		}
		if max < length{
					max = length
				}
		return max;
	}
	//length := 0
	for i :=0;i<len(s);i++{
		ch[s[i]-'A'] = true
	}
	for i :=0;i<26;i++{
		if ch[i] == false{
			continue
		}
		c := 'A' + i
	//	if c == 75{
	//		fmt.Println("break")
	//	}
		idx :=0
		for j:=0;j<len(s);j++{
			if int(s[j])!=c{
				dp[idx] = j
				idx++
			//else{
		//		fmt.Print(j," ")
			}
		}
		//fmt.Println("|",c)
		if idx <= k{
			return len(s)
		}
		if max < dp[k]{
			max = dp[k]
		}
		for j:= k+1;j <idx;j++{
			if dp[j] - dp[j-k-1] - 1 >= max{
				max = dp[j] - dp[j-k-1] - 1
				//fmt.Println(c)
			}
		}
	}
	return max
}
//[[12,13,1,12],[13,4,13,12],[13,8,10,12],[12,13,12,12],[13,13,13,13]]
//"IMNJJTRMJEGMSOLSCCQICIHLQIOGBJAEHQOCRAJQMBIBATGLJDTBNCPIFRDLRIJHRABBJGQAOLIKRLHDRIGERENNMJSDSSMESSTR"
//2 6
//"AABABBA" 1 4
//"KJRGKSKKOKLPADMAGODEDNEBMJMKGAPNLSAKADRLHHDRMJTMFBSIQGHENKABISHQNRFJKEPMFIPNDNEOBRJEHFLIHMDLMCIHLHQN" 5 10
func main() {

	fmt.Println(characterReplacement("KJRGKSKKOKLPADMAGODEDNEBMJMKGAPNLSAKADRLHHDRMJTMFBSIQGHENKABISHQNRFJKEPMFIPNDNEOBRJEHFLIHMDLMCIHLHQN" ,5))
}
