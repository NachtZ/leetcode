package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
	//"math/rand"
	//"sort"
	"regexp"
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
	fmt.Println(regexp.QuoteMeta(""))
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
		fmt.Println("nil")
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
type StrSlice []string

func (s StrSlice) Less(i, j int) bool {
	if len(s[i]) != len(s[j]) {
		return len(s[i]) < len(s[j])
	}
	return s[i] < s[j]
}
func (s StrSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s StrSlice) Len() int {
	return len(s)
}

/***************************************************/

type point struct {
	x int
	y int
	h int
}

type Points []point

func (p Points) Len() int {
	return len(p)
}
func (p Points) Less(i, j int) bool {
	return p[i].h < p[j].h
}

func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func check(x, y int, f [][]int) bool {
	if x < 0 || x >= len(f) || y < 0 || y >= len(f[0]) || f[x][y] == 0 {
		return false
	}
	return true

}

func bfs(f [][]int, x1, y1, x2, y2 int) int {
	movex, movey := []int{1, -1, 0, 0}, []int{0, 0, 1, -1}
	if f[x2][y2] == 0 {
		return -1
	}
	//ret := -1
	v := [50][50]bool{}
	//step := 0
	queue := make([]point, 0, 250)
	queue = append(queue, point{x1, y1, 0})
	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			nx, ny := t.x+movex[i], t.y+movey[i]
			if !check(nx, ny, f) || v[nx][ny] {
				continue
			}
			if nx == x2 && ny == y2 {
				return t.h + 1
			}
			v[nx][ny] = true
			queue = append(queue, point{nx, ny, t.h + 1})
		}
		//step++
	}
	return -1
}

func cutOffTree(f [][]int) int {
	points := []point{}
	for i := 0; i < len(f); i++ {
		for j := 0; j < len(f[0]); j++ {
			if f[i][j] > 1 {
				points = append(points, point{i, j, f[i][j]})
			}
		}
	}
	ret := 0
	sort.Sort(Points(points))
	x, y := 0, 0
	for i := 0; i < len(points); i++ {
		f[points[i].x][points[i].y] = 1
		if points[i].x == x && points[i].y == y {
			continue
		}
		t := bfs(f, x, y, points[i].x, points[i].y)
		if t == -1 {
			return -1
		}
		ret += t
		x, y = points[i].x, points[i].y
	}
	return ret
}

/*
[[0,0,0,3528,2256,9394,3153],[8740,1758,6319,3400,4502,7475,6812],[0,0,3079,6312,0,0,0],[6828,0,0,0,0,0,8145],[6964,4631,0,0,0,4811,0],[0,0,0,0,9734,4696,4246],[3413,8887,0,4766,0,0,0],[7739,0,0,2920,0,5321,2250],[3032,0,3015,0,3269,8582,0]]
*/
func main() {
	fmt.Println(cutOffTree([][]int{
		[]int{54581641, 64080174, 24346381, 69107959},
		[]int{86374198, 61363882, 68783324, 79706116},
		[]int{668150, 92178815, 89819108, 94701471},
		[]int{83920491, 22724204, 46281641, 47531096},
		[]int{89078499, 18904913, 25462145, 60813308},
	}))
}
