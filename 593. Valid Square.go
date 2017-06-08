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
func dis(a, b []int) int {
	return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
}
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	maps := make(map[int]int)
	maps[dis(p1, p2)]++
	maps[dis(p1, p3)]++
	maps[dis(p1, p4)]++
	maps[dis(p2, p3)]++
	maps[dis(p2, p4)]++
	maps[dis(p3, p4)]++
	if _, ok := maps[0]; ok {
		return false
	}
	return len(maps) == 2
}

/*
 */
func main() {
	fmt.Println(validSquare([]int{0, 0}, []int{1, 1}, []int{1, 0}, []int{0, 1}))
}
