package main

import (
	"fmt"
	"math"
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
/*
type IntSlice []int

func (s IntSlice)Less(i,j int)bool{
	return s[i]<s[j]
}
func (s IntSlice)Swap(i,j int){
	s[i],s[j] = s[j],s[i]
}
func (s IntSlice)Len()int{
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

*/
/**************************************************************/
type BigHeap []int

func (h BigHeap) Len() int           { return len(h) }
func (h BigHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h BigHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *BigHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *BigHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type SmallHeap []int

func (h SmallHeap) Len() int           { return len(h) }
func (h SmallHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h SmallHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *SmallHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *SmallHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func threeSort(a, b, c int) (int, int, int) {
	if a > b && a > c {
		if b > c {
			return c, b, a
		}
		return b, c, a
	}
	if b > a && b > c {
		if a > c {
			return c, a, b
		}
		return a, c, b
	}
	if a > b {
		return b, a, c
	}
	return a, b, c
}

func medianSlidingWindow(nums []int, k int) []float64 {
	if k == 1 {
		ret := make([]float64, len(nums))
		for i := 0; i < len(nums); i++ {
			ret[i] = float64(nums[i])
		}
		return ret
	}
	tmpLeft, tmpRight := make(BigHeap, 0, k), make(SmallHeap, 0, k)
	left, right := &tmpLeft, &tmpRight
	ret := make([]float64, 0, len(nums))
	for i := 0; i < k; i++ {
		//fmt.Println("left:", *left)
		//fmt.Println("right:", *right)
		if left.Len() == right.Len() {
			if left.Len() == 0 {
				heap.Push(left, nums[i])
			} else {
				l, m, r := heap.Pop(left).(int), nums[i], heap.Pop(right).(int)
				l, m, r = threeSort(l, m, r)
				heap.Push(left, l)
				heap.Push(left, m)
				heap.Push(right, r)
			}
		} else {
			if right.Len() == 0 {
				t := heap.Pop(left).(int)
				if t > nums[i] {
					heap.Push(right, t)
					heap.Push(left, nums[i])
				} else {
					heap.Push(right, nums[i])
					heap.Push(left, t)
				}
				continue
			}
			if left.Len() > right.Len() {
				l, m, r := heap.Pop(left).(int), nums[i], heap.Pop(right).(int)
				l, m, r = threeSort(l, m, r)
				heap.Push(left, l)
				heap.Push(right, m)
				heap.Push(right, r)
			} else {
				l, m, r := heap.Pop(left).(int), nums[i], heap.Pop(right).(int)
				l, m, r = threeSort(l, m, r)
				heap.Push(left, l)
				heap.Push(left, m)
				heap.Push(right, r)
			}
		}
	}

	for i, j := k-1, -1; i < len(nums); i, j = i+1, j+1 {
		if i >= k {
			//remove the nums[j] from the heap

			l, r := heap.Pop(left).(int), heap.Pop(right).(int)
			//fmt.Println("Try to remove ", nums[j], "with l r", l, r)
			if l == nums[j] || r == nums[j] {
				if l == nums[j] {
					heap.Push(right, r)
				} else {
					heap.Push(left, l)
				}
			} else {
				heap.Push(left, l)
				heap.Push(right, r)

				if l > nums[j] {
					for s := 0; s < len(*left); s++ {
						if (*left)[s] == nums[j] {
							//fmt.Println(" left remove:", (*left)[s])
							heap.Remove(left, s)
							break
						}
					}
				} else {
					for s := 0; s < len(*right); s++ {
						if (*right)[s] == nums[j] {
							//fmt.Println("right remove:", (*right)[s])
							heap.Remove(right, s)
							break
						}
					}
				}
			}
			if len(*left) == 0 || len(*right) == 0 {
				var m, r int
				if len(*left) == 0 {
					m, r = nums[i], heap.Pop(right).(int)
					if m > r {
						m, r = r, m
					}
				} else {
					m, r = nums[i], heap.Pop(left).(int)
					if m > r {
						m, r = r, m
					}
				}
				heap.Push(left, m)
				heap.Push(right, r)
			} else {

				l, m, r := threeSort(heap.Pop(left).(int), nums[i], heap.Pop(right).(int))
				heap.Push(left, l)
				heap.Push(right, r)
				if len(*left) > len(*right) {
					heap.Push(right, m)
				} else {
					heap.Push(left, m)
				}
			}
		}
		//fmt.Println("left:", *left)
		//fmt.Println("right:", *right)
		l, r := heap.Pop(left).(int), heap.Pop(right).(int)

		if k%2 == 0 {
			ret = append(ret, float64(l+r)/2.0)
		} else {
			if left.Len() > right.Len() {
				ret = append(ret, float64(l))
			} else {
				ret = append(ret, float64(r))
			}
		}
		heap.Push(left, l)
		heap.Push(right, r)
	}
	return ret
}
func main() {

	fmt.Println(medianSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
