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
func helper(v [][]int, l, x, y int) {
	if x < 0 || x >= len(v) || y < 0 || y >= len(v[0]) {
		return
	}
	if v[x][y] < l+1 {
		return
	}
	v[x][y] = l + 1
}
func updateMatrix(m [][]int) [][]int {
	v := make([][]int, len(m))
	for i := 0; i < len(v); i++ {
		v[i] = make([]int, len(m[0]))
		for j := 0; j < len(m[0]); j++ {
			if m[i][j] == 0 {
				v[i][j] = 0
			} else {
				v[i][j] = 10001
			}
		}
	}
	found := true
	for l := 0; found && l < 1001; l++ {
		found = false
		for i := 0; i < len(m); i++ {
			for j := 0; j < len(m[0]); j++ {
				if v[i][j] != l {
					continue
				}
				found = true
				helper(v, l, i+1, j)
				helper(v, l, i-1, j)
				helper(v, l, i, j+1)
				helper(v, l, i, j-1)
			}
		}
	}
	return v
}

/**/
func main() {
	m := [][]int{[]int{1, 0, 1, 1, 0, 0, 1, 0, 0, 1}, []int{0, 1, 1, 0, 1, 0, 1, 0, 1, 1}, []int{0, 0, 1, 0, 1, 0, 0, 1, 0, 0}, []int{1, 0, 1, 0, 1, 1, 1, 1, 1, 1}, []int{0, 1, 0, 1, 1, 0, 0, 0, 0, 1}, []int{0, 0, 1, 0, 1, 1, 1, 0, 1, 0}, []int{0, 1, 0, 1, 0, 1, 0, 0, 1, 1}, []int{1, 0, 0, 0, 1, 1, 1, 1, 0, 1}, []int{1, 1, 1, 1, 1, 1, 1, 0, 1, 0}, []int{1, 1, 1, 1, 0, 1, 0, 0, 1, 1}}
	for i := 0; i < len(m); i++ {
		fmt.Println(m[i])
	}
	v := updateMatrix(m)
	for i := 0; i < len(v); i++ {
		fmt.Println(v[i])
	}
}
