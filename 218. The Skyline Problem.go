package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
	//"math/rand"
	//"sort"
	//"regexp"
	"container/heap"
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

/**************************************************************/

type IntSlice []int

func (s IntSlice) Less(i, j int) bool {
	return s[i] > s[j]
}
func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s IntSlice) Len() int {
	return len(s)
}

/**************************************************************/
type BdHeap [][]int

func (h BdHeap) Len() int { return len(h) }
func (h BdHeap) Less(i, j int) bool {
	if h[i][0] != h[j][0] {
		return h[i][0] < h[j][0]
	}
	return h[i][1] < h[j][1]
}
func (h BdHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntSlice) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntSlice) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getSkyline(buildings [][]int) [][]int {
	h := new(IntSlice)
	heap.Init(h)
	ret := [][]int{}
	tmpPair := make([][]int, 0, 2*len(buildings))
	for _, t := range buildings {
		tmpPair = append(tmpPair, []int{t[0], -t[2]})
		tmpPair = append(tmpPair, []int{t[1], t[2]})
	}
	sort.Sort(BdHeap(tmpPair))
	//fmt.Println(tmpPair)
	pre := 0
	heap.Push(h, 0)
	for i := 0; i < len(tmpPair); i++ {
		if tmpPair[i][1] < 0 {
			heap.Push(h, -tmpPair[i][1])
		} else {
			idx := 0
			for idx = 0; idx < len(*h); idx++ {
				if (*h)[idx] == tmpPair[i][1] {
					break
				}
			}
			//fmt.Println("remove idx", idx, tmpPair[i][1], (*h)[idx])
			heap.Remove(h, idx)
		}
		//fmt.Println("heap:", *h)
		cur := (*h)[0]
		if pre != cur {
			ret = append(ret, []int{tmpPair[i][0], cur})
			pre = cur
		}
		//fmt.Println("ret", ret)
	}
	return ret
}

func main() {

	fmt.Println(getSkyline([][]int{[]int{0, 2, 3}, []int{2, 5, 3}}))
}
