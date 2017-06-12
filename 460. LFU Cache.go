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

/*************************************************************/
type element struct {
	val    int
	key    int
	fre    *freq
	next   *element
	before *element
}

type freq struct {
	next   *freq
	before *freq
	time   int
	eleH   *element //element linklist head
	eleE   *element //element linklist end
}

type LFUCache struct {
	dict  map[int]*element
	freqs map[int]*freq //map of freqence
	freqH freq          // freqence link list, from small to large
	cap   int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		make(map[int]*element),
		make(map[int]*freq),
		freq{nil, nil, 0, nil, nil},
		capacity,
	}
}

func (this *LFUCache) releaseLink(f *freq) {
	if f == nil {
		return
	}
	if f.next == nil {
		f.before.next = nil
		return
	}
	b, n := f.before, f.next
	f.before.next = n
	f.next.before = b
	delete(this.freqs, f.time)
}

func (this *LFUCache) releaseEle(f *freq, v *element) {
	b, n := v.before, v.next
	if b == nil && n == nil {
		//free tfre
		this.releaseLink(f)
	} else {
		if n != nil {
			n.before = b
		} else {
			f.eleE = b
		}
		if b != nil {
			b.next = n
		} else {
			f.eleH = n
		}
	}
	v.before, v.next = nil, nil

}

func (this *LFUCache) Get(key int) int {
	v, ok := this.dict[key]
	if !ok {
		return -1
	}
	tfreq := v.fre
	if tfreq.next != nil && tfreq.time+1 == tfreq.next.time {
		if tfreq.eleH == tfreq.eleE {
			tfreq.next.before, tfreq.before.next = tfreq.before, tfreq.next
			delete(this.freqs, tfreq.time)

		} else {
			this.releaseEle(tfreq, v)
		}
		tfreq.next.eleE.next = v
		v.before = tfreq.next.eleE
		tfreq.next.eleE = v
		v.fre = tfreq.next
	} else {
		if tfreq.eleH == tfreq.eleE {
			delete(this.freqs, tfreq.time)
			tfreq.time++
			this.freqs[tfreq.time] = tfreq
		} else {
			this.releaseEle(tfreq, v)
			tmp := &freq{tfreq.next, tfreq, tfreq.time + 1, v, v}
			tfreq.next = tmp
			if tmp.next != nil {
				tmp.next.before = tmp
			}
			this.freqs[tmp.time] = tmp
			v.fre = tmp
		}
	}
	/*{
		tmp := &freq{tfreq.next, tfreq, tfreq.time + 1, v, v}
		tfreq.next = tmp
		v.fre = tmp
		if tmp.next != nil {
			tmp.next.before = tmp
		}
		this.releaseEle(tfreq, v)
		v.before, v.next = nil, nil
	}*/
	//fmt.Println("NExt", this.freqH.next.time)
	return v.val
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}
	if e, ok := this.dict[key]; ok {
		e.val = value
		//		this.print()
		this.Get(key)
		//		this.print()
		return
	}
	if len(this.dict) >= this.cap {
		f := this.freqH.next
		e := f.eleH
		e.fre = f
		if f.eleH == f.eleE {
			this.freqH.next = f.next
			if f.next != nil {
				this.freqH.next.before = &this.freqH
			}
			delete(this.freqs, f.time)
		} else {
			f.eleH = f.eleH.next
			f.eleH.before = nil
		}
		delete(this.dict, e.key)
	}
	ele := &element{value, key, nil, nil, nil}
	this.dict[key] = ele
	f, ok := this.freqs[1]
	if ok {
		f.eleE.next, ele.before = ele, f.eleE
		f.eleE = ele
		ele.fre = f
	} else {
		f = &freq{this.freqH.next, &this.freqH, 1, ele, ele}
		if f.next != nil {
			f.next.before = f
		}
		this.freqH.next = f
		this.freqs[1] = f
		ele.fre = f

	}
}
func (this *LFUCache) print() {
	fmt.Println("---")
	t := this.freqH.next
	for t != nil {
		fmt.Printf("%p ", t)
		fmt.Print(t, ": ")
		tmp := t.eleH
		for tmp != nil {
			fmt.Print(tmp.key, ":", tmp.val, " ")
			tmp = tmp.next
		}
		t = t.next
	}
	fmt.Println("\n---")
}

/*
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	fmt.Println(c.Get(1))
	c.Put(3, 3)
	fmt.Println(c.Get(2))
	fmt.Println(c.Get(3))
	c.Put(4, 4)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(3))
	fmt.Println(c.Get(4))
*/
/*
	c := Constructor(0)
	c.Put(0, 0)
	fmt.Println(c.Get(0))
*/
/*
["LFUCache","put","put","get","put","get","get","put","get","get","get"]
[[2],[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]]
["LFUCache","put","get"]
[[0],[0,0],[0]]
["LFUCache","put","get","put","get","get"]
[[1],[2,1],[2],[3,2],[2],[3]]
["LFUCache","get","put","get","put","put","get","get"]
[[2],[2],[2,6],[1],[1,5],[1,2],[1],[2]]
["LFUCache","put","put","put","put","get"]
[[2],[3,1],[2,1],[2,2],[4,4],[2]]
*/
/*
	c := Constructor(2)
	fmt.Println(c.Get(2))
	c.Put(2, 6)
	c.print()
	fmt.Println(c.Get(1))
	c.print()
	c.Put(1, 5)
	c.print()
	c.Put(1, 2)
	c.print()
	fmt.Println(c.Get(1))
	c.print()
	fmt.Println(c.Get(2))
	c.print()
*/
/*
	c := Constructor(2)

	c.Put(3, 1)
	c.print()

	c.Put(2, 1)
	c.print()
	c.Put(2, 2)
	c.print()
	c.Put(4, 4)
	c.print()
	fmt.Println(c.Get(2))
	c.print()
*/
func main() {
	c := Constructor(1)

	c.Put(2, 1)
	c.print()
	fmt.Println(c.Get(2))
	c.print()
	c.Put(3, 2)
	c.print()
	fmt.Println(c.Get(2))
	c.print()
	fmt.Println(c.Get(3))

	c.print()
}
