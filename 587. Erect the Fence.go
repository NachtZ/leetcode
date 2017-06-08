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
type Points []Point

var maxd float64

func (p Points) Less(i, j int) bool {
	if p[i].Y != p[j].Y {
		return p[i].Y < p[j].Y
	}
	return p[i].X < p[j].X
}

func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Points) Len() int {
	return len(p)
}

type spoint struct {
	p     Point
	Radio float64
	dis   int
}

type sPoints []spoint

func (p sPoints) Less(i, j int) bool {
	if p[i].Radio == p[j].Radio {
		if p[i].Radio < math.Pi/2 {
			return p[i].dis < p[j].dis
		} else if p[i].Radio > math.Pi/2 {
			return p[i].dis > p[j].dis
		} else if maxd <= math.Pi/2 {
			return p[i].dis > p[j].dis
		} else {
			return p[i].dis < p[j].dis
		}
	}
	return p[i].Radio < p[j].Radio
}

func (p sPoints) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p sPoints) Len() int {
	return len(p)
}

func cross(a, b, c Point) bool {
	return (b.X-a.X)*(c.Y-b.Y)-(b.Y-a.Y)*(c.X-b.X) >= 0
}

func outerTrees(points []Point) []Point {
	if len(points) < 4 {
		return points
	}
	sort.Sort(Points(points))
	ret := make([]Point, 0)
	tmpPoint := make(sPoints, len(points)-1)
	maxd = 0
	for i := 1; i < len(points); i++ {
		radio := math.Pi / 2
		if points[i].X != points[0].X {

			radio = math.Atan(float64(points[i].Y-points[0].Y) / float64(points[i].X-points[0].X))
			if radio < 0 {
				radio += math.Pi * 2
			}
			fmt.Println(radio, float64(points[i].Y-points[0].Y)/float64(points[i].X-points[0].X))
		}
		if maxd < radio {
			maxd = radio
		}
		tmpPoint[i-1].Radio = radio
		tmpPoint[i-1].p = points[i]
		tmpPoint[i-1].dis = (points[i].Y-points[0].Y)*(points[i].Y-points[0].Y) + (points[i].X-points[0].X)*(points[i].X-points[0].X)
	}
	sort.Sort(tmpPoint)
	for i := 0; i < len(tmpPoint); i++ {
		fmt.Println(tmpPoint[i].Radio, tmpPoint[i].p)
	}
	ret = append(ret, points[0])
	ret = append(ret, tmpPoint[0].p)
	for i := 1; i < len(tmpPoint); i++ {
		t := tmpPoint[i].p
		flag := true
		for len(ret) > 2 && flag {
			t1, t2 := ret[len(ret)-1], ret[len(ret)-2]
			if cross(t2, t1, t) {
				flag = false
			} else {
				ret = ret[0 : len(ret)-1]
			}
		}
		ret = append(ret, t)
	}
	fmt.Println(maxd, math.Pi/2)
	return ret
}

/*
 */
func main() {
	t := []int{0, 0, 0, 1, 0, 2, 1, 2, 2, 2, 3, 2, 3, 1, 3, 0, 2, 0}
	p := []Point{}
	for i := 0; i < len(t); i += 2 {
		p = append(p, Point{t[i], t[i+1]})
	}
	fmt.Println(outerTrees(p))
}
