package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
	//"math/rand"
	//"sort"

	"sort"
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

type Interval struct {
	Start int
	End   int
}

func useLib() {
	fmt.Println(strconv.Itoa(1))
	fmt.Println(strings.Compare("1", "2"))
	fmt.Println(math.Abs(1.0))
	fmt.Println(unicode.IsDigit('1'))
	fmt.Println(sort.IsSorted(nil))
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

func printTreeNode(root *TreeNode) {
	if root == nil {
		return
	}
	left, right := "nil", "nil"
	if root.Left != nil {
		left = strconv.Itoa(root.Left.Val)
	}
	if root.Right != nil {
		right = strconv.Itoa(root.Right.Val)
	}
	fmt.Println(root.Val, ":", left, right)
	printTreeNode(root.Left)
	printTreeNode(root.Right)
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
func printList(root *ListNode) {
	for root != nil {
		fmt.Print(root.Val, "->")
		root = root.Next
	}
	fmt.Println("nil")
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type trie struct {
	ch    byte
	child [26]*trie
	flag  bool // true means here located a word.
}

func (this *trie) insertWords(str string) {
	ptr := this
	for idx, s := range str {
		if ptr.child[s-'a'] == nil {
			ptr.child[s-'a'] = new(trie)
			ptr.child[s-'a'].ch = str[idx]
		}
		if idx == len(str)-1 {
			ptr.child[s-'a'].flag = true
		}
		ptr = ptr.child[s-'a']
	}
}

func (this *trie) searchWords(str string) *trie {
	ptr := this
	for idx, s := range str {
		if ptr.child[s-'a'] == nil {
			return nil
		}
		if idx == len(str)-1 {
			if ptr.child[s-'a'].flag == true {
				return ptr.child[s-'a']
			}
			return nil
		}
		ptr = ptr.child[s-'a']

	}
	return nil
}

func buildTrie(words []string) *trie {
	root := &trie{}
	for _, w := range words {
		root.insertWords(w)
	}
	return root
}

/**************************************************************/

type IntSlice []int

func (s IntSlice) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s IntSlice) Len() int {
	return len(s)
}

type StrSlice []string

func (s StrSlice) Less(i, j int) bool {
	if len(s[i]) != len(s[j]) {
		return len(s[i]) > len(s[j])
	}
	return s[i] < s[j]
}
func (s StrSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s StrSlice) Len() int {
	return len(s)
}

/**************************************************************/

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

/**************************************************************/
func getPalidromi(n, normal string) string {
	v, _ := strconv.Atoi(n)
	vn, _ := strconv.Atoi(normal)
	n1, n2 := strconv.Itoa(v+1), strconv.Itoa(v-1)
	tmp := 1
	if len(n1)*2 != len(normal) {
		tmp = 2
	}
	for i := len(n1) - tmp; i >= 0; i-- {
		n1 = n1 + string(n1[i])
	}
	tmp = 1
	if len(n2)*2 != len(normal) {
		tmp = 2
	}
	for i := len(n2) - tmp; i >= 0; i-- {
		n2 = n2 + string(n2[i])
	}
	v1, _ := strconv.Atoi(n1)
	v2, _ := strconv.Atoi(n2)
	if abs(v2-vn) <= abs(v1-vn) {
		return n2
	}
	return n1
}

func isPalindromi(n string) bool {
	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		if n[i] != n[j] {
			return false
		}
	}
	return true
}
func nearestPalindromic(n string) string {
	v, _ := strconv.Atoi(n)

	if len(n) == 1 {
		return strconv.Itoa(v - 1)
	}
	tail := ""
	for i := 0; i < len(n)/2; i++ {
		tail = string(n[i]) + tail
	}
	ret := ""
	if len(n)%2 == 0 {
		if isPalindromi(n) {
			ret = getPalidromi(n[:len(n)/2], n)
		} else {
			ret = getPalidromi(n[:len(n)/2], n)
			ret1 := n[0:len(n)/2] + tail
			v1, _ := strconv.Atoi(ret)
			v2, _ := strconv.Atoi(ret1)
			if abs(v1-v) == abs(v2-v) {
				if v2 < v1 {
					ret = ret1
				}
			} else {
				if abs(v2-v) < abs(v1-v) {
					ret = ret1
				}
			}
		}
	} else {
		if isPalindromi(n) {
			ret = getPalidromi(n[:len(n)/2+1], n)
		} else {
			ret = getPalidromi(n[:len(n)/2+1], n)
			ret1 := n[0:len(n)/2+1] + tail
			v1, _ := strconv.Atoi(ret)
			v2, _ := strconv.Atoi(ret1)
			if abs(v1-v) == abs(v2-v) {
				if v2 < v1 {
					ret = ret1
				}
			} else {
				if abs(v2-v) < abs(v1-v) {
					ret = ret1
				}
			}
		}
	}
	m := ""
	for i := 1; i < len(n); i++ {
		m += "9"
	}
	maxv := "1"
	for i := 1; i < len(n); i++ {
		maxv += "0"
	}
	maxv += "1"
	vm, _ := strconv.Atoi(m)
	vmaxv, _ := strconv.Atoi(maxv)
	vr, _ := strconv.Atoi(ret)
	if abs(v-vm) <= abs(vr-v) && abs(v-vm) <= abs(vmaxv-v) {
		return m
	}
	if abs(vmaxv-v) < abs(vr-v) {
		return maxv
	}
	return ret

}

/*
"1"
"2"
"10"
"100"
"111"
"1111"
"11111"
"111111"
"55555"
"555555"
"99"
"100"
"121"
"123"
*/

func main() {
	fmt.Println(nearestPalindromic("1283"))
}
