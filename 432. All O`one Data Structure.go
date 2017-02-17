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
*/
/**************************************************************/

type Node struct {
	Before *Node
	After  *Node // After is larger than before
	Value  int
	Key    string
}

type AllOne struct {
	Map    map[string]*Node
	MaxEnd *Node
	MinEnd *Node
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	var ret AllOne
	ret.Map = make(map[string]*Node)
	return ret
}

func (this *AllOne) Print() {
	fmt.Println("Max:", this.MaxEnd, "Min:", this.MinEnd)
	for tmp := this.MinEnd; tmp != nil; tmp = tmp.After {
		fmt.Print(tmp.Key, tmp.Value, "->")
	}
	fmt.Println()
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	if node, ok := this.Map[key]; ok {
		node.Value++
		if node.After != nil && node.After.Value < node.Value { // change two value
			var tmp *Node
			for tmp = node.After; tmp.After != nil && tmp.Value == tmp.After.Value; tmp = tmp.After {

			}
			tmp.Value, node.Value = node.Value, tmp.Value
			tmp.Key, node.Key = node.Key, tmp.Key
			this.Map[key], this.Map[node.Key] = tmp, node
		}
	} else {
		tmp := &Node{
			nil,
			this.MinEnd,
			1,
			key,
		}
		if this.MinEnd != nil {
			this.MinEnd.Before = tmp
		}
		this.MinEnd = tmp
		if this.MaxEnd == nil {
			this.MaxEnd = tmp
		}
		this.Map[key] = tmp
	}
	//	fmt.Println(key, this.Map[key].Value)
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	node, ok := this.Map[key]
	if ok == false {
		return
	}
	node.Value--
	//fmt.Println(key, this.Map[key].Value)
	if node.Value == 0 { // delete the node
		if this.MaxEnd == node {
			this.MaxEnd = node.Before
		}
		delete(this.Map, key)
		if this.MinEnd == node {
			this.MinEnd = node.After
			if this.MinEnd != nil {
				this.MinEnd.Before = nil
			}
		} else {
			if node.After != nil {
				node.After.Before = node.Before
			}
			if node.Before != nil {
				node.Before.After = node.After
			}
		}

	} else {

		if node.Before != nil && node.Before.Value > node.Value {
			var tmp *Node
			for tmp = node.Before; tmp.Before != nil && tmp.Before.Value == tmp.Value; tmp = tmp.Before {

			}
			tmp.Value, node.Value = node.Value, tmp.Value
			tmp.Key, node.Key = node.Key, tmp.Key
			this.Map[key], this.Map[node.Key] = tmp, node
		}
	}

}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	if this.MaxEnd != nil {
		return this.MaxEnd.Key
	}
	return ""
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	if this.MinEnd != nil {
		return this.MinEnd.Key
	}
	return ""
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */

func main() {
	//fmt.Println(findContentChildren([]int{1,2},[]int{1,2,3}))
	//fmt.Println(pathSum(buildTree([]int{1,2}),2))
	//fmt.Println(findKthNumber(13,2))
	obj := Constructor()
	obj.Inc("hello")
	obj.Inc("world")
	obj.Inc("leet")
	obj.Inc("code")
	obj.Inc("DS")
	obj.Inc("leet")
	obj.Print()

	fmt.Println("Max:", obj.GetMaxKey())
	obj.Inc("DS")
	obj.Print()

	obj.Dec("leet")
	obj.Print()
	fmt.Println("Max:", obj.GetMaxKey())
	obj.Dec("DS")
	obj.Inc("hello")
	obj.Print()
	fmt.Println("Max:", obj.GetMaxKey())
	obj.Inc("hello")
	obj.Inc("hello")
	obj.Dec("world")
	obj.Dec("leet")
	obj.Dec("code")
	obj.Dec("DS")
	obj.Print()
	fmt.Println("Max:", obj.GetMaxKey())
	obj.Inc("new")
	obj.Inc("new")
	obj.Inc("new")
	obj.Inc("new")
	obj.Inc("new")
	obj.Inc("new")
	obj.Print()
	fmt.Println("Max:", obj.GetMaxKey())
	fmt.Println("Min:", obj.GetMinKey())

}
